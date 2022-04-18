package simulate_tx_api

import "github.com/ThreeAndTwo/simulate-tx-api/providers"

type ISimulate interface {
	AddForkEnv(chainId, name string) error
	RenameForkEnv(forkId, chainId, name string) error
	DeleteForkEnv(forkId string) error
	SimulateTxForFork(forkId, params string) (string, error)
	Simulate(params string) (string, error)
}

type Simulate struct {
	Account string
	Project string
	Token   string
	Tps     int
}

type Platform string

const (
	SimulateTenderly    Platform = "tenderly"
	SimulateBlockNative Platform = "blocknative"
)

func NewSimulate(account, project, token string, tps int) *Simulate {
	return &Simulate{Account: account, Project: project, Token: token, Tps: tps}
}

func (s *Simulate) SimulateGetter(platform Platform) ISimulate {
	switch platform {
	case SimulateTenderly:
		return providers.NewTenderSimulate(s.Account, s.Project, s.Token, s.Tps)
	case SimulateBlockNative:
		return nil
	default:
		return providers.NewTenderSimulate(s.Account, s.Project, s.Token, s.Tps)
	}
}
