package simulate_tx_api

import (
	"fmt"
	"github.com/ThreeAndTwo/simulate-tx-api/providers"
	"os"
	"testing"
)

type actionTy int

const (
	actionAddProject actionTy = iota
	actionRenameProject
	actionSimulate
	actionAddEnv
	actionRenameEnv
	actionDeleteEnv
	actionForkSimulate
	actionJsonRpc
)

type testSimulate struct {
	name    string
	action  actionTy
	account string
	token   string
	project string
	tps     int
	params  string
}

func TestTenderly(t *testing.T) {
	tests := []testSimulate{
		{
			name:    "add project on tenderly",
			action:  actionAddProject,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "project",
			tps:     1,
			params:  "",
		},
		{
			name:    "rename project on tenderly",
			action:  actionRenameProject,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "project",
			tps:     1,
			params:  "",
		},
		{
			name:    "simulate failed for tenderly",
			action:  actionSimulate,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     1,
			params:  `{"network_id":"1","block_number":14365389,"transaction_index":0,"from":"0x7da5eacc8628f22d5e56ed0018751a8921942e38","input":"0xf8ca0f85174876e80083030d40947cad06b811b5d9d3ff197c1a046abcbc0efbcbc980b864d675fd260000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c626a0189c5d42fea496685df9488e8c2b761a278a7a6744383a9a6d7927dc11f70c29a02a948c7f03f80a30cf7fe5262c0f95068fa2c111b655d6a73cc223b151b6c0dc","to":"0x7cad06b811b5d9d3ff197c1a046abcbc0efbcbc9","gas":200000,"gas_price":"100000000000","value":"0","access_list":[],"generate_access_list":true,"save":false,"source":"dashboard"}`,
		},
		{
			name:    "simulate success for tenderly",
			action:  actionSimulate,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     1,
			params:  `{"network_id":"1","block_number":14365440,"transaction_index":0,"from":"0x7da5eacc8628f22d5e56ed0018751a8921942e38","input":"0xd675fd260000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c7","to":"0x7cad06b811b5d9d3ff197c1a046abcbc0efbcbc9","gas":200000,"gas_price":"100000000000","value":"0","access_list":[],"generate_access_list":true,"save":false,"source":"dashboard"}`,
		},
		{
			name:    "params error",
			action:  actionSimulate,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     0,
			params:  "",
		},
		{
			name:    "test add fork env",
			action:  actionAddEnv,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     1,
		},
		{
			name:    "test add fork env, project error",
			action:  actionAddEnv,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "test-fork",
			tps:     1,
		},
		{
			name:    "test rename fork env",
			action:  actionRenameEnv,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     1,
		},
		{
			name:    "test delete fork env",
			action:  actionDeleteEnv,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     1,
		},
		{
			name:    "test simulation fork env",
			action:  actionForkSimulate,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv("TOKEN"),
			project: "aaa-bbb",
			tps:     1,
		},
		{
			name:    "test request json rpc",
			action:  actionJsonRpc,
			account: os.Getenv("ACCOUNT"),
			token:   os.Getenv(""),
			project: "",
			tps:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulator := NewSimulate(tt.account, tt.project, tt.token, tt.tps).SimulateGetter(SimulateTenderly)
			switch tt.action {
			case actionAddProject:
				tt.testAddProject(simulator)
			case actionRenameProject:
				tt.testRenameProject(simulator)
			case actionSimulate:
				tt.testSimulate(simulator)
			case actionAddEnv:
				tt.testAddEnv(simulator)
			case actionRenameEnv:
				tt.testRenameEnv(simulator)
			case actionDeleteEnv:
				tt.testDeleteEnv(simulator)
			case actionForkSimulate:
				tt.testForkSimulation(simulator)
			case actionJsonRpc:
				tt.testReqJsonRpc(simulator)
			}
		})
	}
}

func (ts *testSimulate) testReqJsonRpc(simulator ISimulate) {
	var p1 []interface{}
	account := []string{"0x8b4941915D7E2971E583976c66Da3a84A6E1936b"}
	p1 = append(p1, account)
	p1 = append(p1, "0x3e8")

	var p2 []interface{}
	p2 = append(p2, []string{})
	p2 = append(p2, "")

	var tests = []struct {
		name   string
		rpc    string
		params *providers.RpcParams
	}{
		{
			name: "normal",
			rpc:  "https://rpc.tenderly.co/fork/" + os.Getenv("FORKID"),
			params: &providers.RpcParams{
				Method: "tenderly_addBalance",
				Params: p1,
			},
		},
		{
			name: "account is null",
			rpc:  "https://rpc.tenderly.co/fork/" + os.Getenv("FORKID"),
			params: &providers.RpcParams{
				Method: "tenderly_addBalance",
				Params: p2,
			},
		},
	}

	for _, tt := range tests {
		res, err := simulator.ReqJsonRpc(tt.rpc, tt.params)
		if err != nil {
			_ = fmt.Errorf("req json rpc for %s error: %s \n", ts.name, err)
			return
		}
		fmt.Printf("req json rpc for %s, result: %s \n", ts.name, res)
	}
}

func (ts *testSimulate) testAddProject(simulator ISimulate) {
	tests := []struct {
		name    string
		proName string
	}{
		{
			name:    "test add project",
			proName: "test-project",
		},
		{
			name:    "project name is null",
			proName: "",
		},
	}

	for _, tt := range tests {
		res, err := simulator.AddProject(tt.proName)
		if err != nil {
			_ = fmt.Errorf("add env for %s error: %s \n", ts.name, err)
			return
		}
		fmt.Printf("add env for %s, result: %s \n", ts.name, res)
	}
}

func (ts *testSimulate) testRenameProject(simulator ISimulate) {
	tests := []struct {
		name    string
		proName string
	}{
		{
			name:    "test add project",
			proName: "Project",
		},
		{
			name:    "project name is null",
			proName: "",
		},
	}

	for _, tt := range tests {
		res, err := simulator.RenameProject(tt.proName)
		if err != nil {
			_ = fmt.Errorf("add env for %s error: %s \n", ts.name, err)
			return
		}
		fmt.Printf("add env for %s, result: %s \n", ts.name, res)
	}
}

func (ts *testSimulate) testSimulate(simulator ISimulate) {
	res, err := simulator.Simulate(ts.params)
	if err != nil {
		_ = fmt.Errorf("add env for %s error: %s \n", ts.name, err)
		return
	}
	fmt.Printf("add env for %s, result: %s \n", ts.name, res)
}

func (ts *testSimulate) testAddEnv(simulator ISimulate) {
	tests := []struct {
		name    string
		chainId string
		envName string
	}{
		{
			name:    "test ethereum mainnet",
			chainId: "1",
			envName: "aaa_ethereum",
		},
		{
			name:    "test bsc mainnet",
			chainId: "56",
			envName: "aaa_bsc",
		},
		{
			name:    "test wrong chainId",
			chainId: "0",
			envName: "aaa_err_chainId",
		},
		{
			name:    "one char for name",
			chainId: "1",
			envName: "aaa_1_char",
		},
	}

	for _, tt := range tests {
		res, err := simulator.AddForkEnv(tt.chainId, tt.envName)
		if err != nil {
			_ = fmt.Errorf("add env for %s error: %s \n", tt.name, err)
			continue
		}
		fmt.Printf("add env for %s, result:%s \n", tt.name, res)
	}
}

func (ts *testSimulate) testRenameEnv(simulator ISimulate) {
	tests := []struct {
		name    string
		forkId  string
		chainId string
		envName string
	}{
		{
			name:    "test ethereum mainnet",
			forkId:  "5051ee85-1441-47ac-a213-f5054ddcba24",
			chainId: "1",
			envName: "ethereum",
		},
		{
			name:    "test bsc mainnet",
			forkId:  "b7078201-445b-4c09-b93b-d52f8066bb4f",
			chainId: "56",
			envName: "bsc",
		},
		{
			name:    "test wrong chainId",
			forkId:  "",
			chainId: "0",
			envName: "aaa_err_chainId",
		},
		{
			name:    "one char for name",
			forkId:  "c6788c5d-1e12-469b-b1a1-7ec8152a6ad0",
			chainId: "1",
			envName: "1",
		},
	}

	for _, tt := range tests {
		res, err := simulator.RenameForkEnv(tt.forkId, tt.chainId, tt.envName)
		if err != nil {
			_ = fmt.Errorf("rename env for %s error: %s \n", tt.name, err)
			continue
		}
		fmt.Printf("rename env for %s, result: %s", tt.name, res)
	}
}

func (ts *testSimulate) testDeleteEnv(simulator ISimulate) {
	test := []struct {
		name   string
		forkId string
	}{
		{
			name:   "normal forkId",
			forkId: "c6788c5d-1e12-469b-b1a1-7ec8152a6ad0",
		},
		{
			name:   "not exists forkId",
			forkId: "8a5d135f-447b-4fa3-9ed2-bcb91498d39a",
		},
	}

	for _, tt := range test {
		res, err := simulator.DeleteForkEnv(tt.forkId)
		if err != nil {
			_ = fmt.Errorf("delete env for %s error: %s \n", tt.name, err)
			continue
		}
		fmt.Printf("delete env for %s, result: %s", tt.name, res)
	}
}

func (ts *testSimulate) testForkSimulation(simulator ISimulate) {
	tests := []struct {
		name   string
		forkId string
		params string
	}{
		{
			name:   "normal",
			forkId: "5051ee85-1441-47ac-a213-f5054ddcba24",
			params: "{\"network_id\":\"1\",\"block_number\":14365440,\"transaction_index\":0,\"from\":\"0x7da5eacc8628f22d5e56ed0018751a8921942e38\",\"input\":\"0xd675fd260000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c7\",\"to\":\"0x7cad06b811b5d9d3ff197c1a046abcbc0efbcbc9\",\"gas\":200000,\"gas_price\":\"100000000000\",\"value\":\"0\",\"access_list\":[],\"generate_access_list\":true,\"save\":true,\"source\":\"dashboard\"}",
		},
		{
			name:   "failed",
			forkId: "",
			params: "",
		},
	}

	for _, tt := range tests {
		res, err := simulator.SimulateTxForFork(tt.forkId, tt.params)
		if err != nil {
			_ = fmt.Errorf("fork simulation env for %s error: %s \n", tt.name, err)
			continue
		}
		fmt.Printf("simulate tx for fork env for %s res:%s \n", tt.name, res)
	}
}
