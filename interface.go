package simulate_tx_api

import "github.com/ThreeAndTwo/simulate-tx-api/providers"

type ISimulate interface {
	AddProject(name string) (string, error)
	RenameProject(name string) (string, error)
	AddForkEnv(chainId, name string) (string, error)
	RenameForkEnv(forkId, chainId, name string) (string, error)
	DeleteForkEnv(forkId string) (string, error)
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
		return providers.NewTenderly(s.Account, s.Project, s.Token, s.Tps)
	case SimulateBlockNative:
		return nil
	default:
		return providers.NewTenderly(s.Account, s.Project, s.Token, s.Tps)
	}
}
