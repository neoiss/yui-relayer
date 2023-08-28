package module

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/neoiss/yui-relayer/config"
	"github.com/neoiss/yui-relayer/provers/mock"
	"github.com/spf13/cobra"
)

type Module struct{}

var _ config.ModuleI = (*Module)(nil)

// Name returns the name of the module
func (Module) Name() string {
	return "mock-client"
}

// RegisterInterfaces register the module interfaces to protobuf Any.
func (Module) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	mock.RegisterInterfaces(registry)
}

// GetCmd returns the command
func (Module) GetCmd(ctx *config.Context) *cobra.Command {
	return nil
}
