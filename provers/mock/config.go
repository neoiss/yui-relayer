package mock

import (
	"github.com/neoiss/yui-relayer/core"
)

var _ core.ProverConfig = (*ProverConfig)(nil)

func (c *ProverConfig) Build(chain core.Chain) (core.Prover, error) {
	return NewProver(chain), nil
}
