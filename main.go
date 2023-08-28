package main

import (
	"log"

	tendermint "github.com/neoiss/yui-relayer/chains/tendermint/module"
	"github.com/neoiss/yui-relayer/cmd"
	mock "github.com/neoiss/yui-relayer/provers/mock/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
