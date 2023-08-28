package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/neoiss/yui-relayer/config"
	"github.com/neoiss/yui-relayer/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func chainsCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chains",
		Short: "manage chain configurations",
	}

	cmd.AddCommand(
		chainsAddDirCmd(ctx),
	)

	return cmd
}

func chainsAddDirCmd(ctx *config.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "add-dir [dir]",
		Args: cobra.ExactArgs(1),
		Short: `Add new chains to the configuration file from a directory 
		full of chain configuration, useful for adding testnet configurations`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if err := filesAdd(ctx, args[0]); err != nil {
				return err
			}
			return overWriteConfig(ctx, cmd)
		},
	}

	return cmd
}

func filesAdd(ctx *config.Context, dir string) error {
	dir = path.Clean(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		pth := fmt.Sprintf("%s/%s", dir, f.Name())
		if f.IsDir() {
			fmt.Printf("directory at %s, skipping...\n", pth)
			continue
		}
		byt, err := ioutil.ReadFile(pth)
		if err != nil {
			fmt.Printf("failed to read file %s, skipping...\n", pth)
			continue
		}
		var c core.ChainProverConfig
		if err := json.Unmarshal(byt, &c); err != nil {
			fmt.Printf("failed to unmarshal file %s, skipping...\n", pth)
			continue
		}
		if err := c.Init(ctx.Codec); err != nil {
			return err
		}
		if err = ctx.Config.AddChain(ctx.Codec, c); err != nil {
			fmt.Printf("%s: %s\n", pth, err.Error())
			continue
		}
		chain, err := c.Build()
		if err != nil {
			return err
		}
		fmt.Printf("added %s...\n", chain.ChainID())
	}
	return nil
}

func overWriteConfig(ctx *config.Context, cmd *cobra.Command) error {
	home, err := cmd.Flags().GetString(flags.FlagHome)
	if err != nil {
		return err
	}

	cfgPath := path.Join(home, "config", "config.yaml")
	if _, err = os.Stat(cfgPath); err == nil {
		viper.SetConfigFile(cfgPath)
		if err = viper.ReadInConfig(); err == nil {
			// ensure validateConfig runs properly
			err = config.InitChains(ctx, homePath, debug)
			if err != nil {
				return err
			}

			// marshal the new config
			out, err := config.MarshalJSON(*ctx.Config)
			if err != nil {
				return err
			}

			// overwrite the config file
			err = ioutil.WriteFile(viper.ConfigFileUsed(), out, 0600)
			if err != nil {
				return err
			}
		}
	}
	return err
}
