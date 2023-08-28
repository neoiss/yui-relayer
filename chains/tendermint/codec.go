package tendermint

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/neoiss/yui-relayer/core"
)

// RegisterInterfaces register the module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfig)(nil),
		&ChainConfig{},
	)
	registry.RegisterImplementations(
		(*core.ProverConfig)(nil),
		&ProverConfig{},
	)
}
