package simulate_tx_api

import "github.com/ThreeAndTwo/simulate-tx-api/providers"

type ISimulate interface {
	Simulate(params string) (string, error)
}

type Simulate struct {
	Account string
	Token   string
	Tps     int
}

type Platform string

const (
	SimulateTenderly    Platform = "tenderly"
	SimulateBlockNative Platform = "blocknative"
)

func NewSimulate(account string, token string, tps int) *Simulate {
	return &Simulate{Account: account, Token: token, Tps: tps}
}

func (s *Simulate) SimulateGetter(platform Platform) ISimulate {
	switch platform {
	case SimulateTenderly:
		return providers.NewTenderSimulate(s.Account, s.Token, s.Tps)
	default:
		return providers.NewTenderSimulate(s.Account, s.Token, s.Tps)
	}
}
