package simulate_tx_api

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeAndTwo/simulate-tx-api/providers"
	"os"
	"testing"
)

func TestTenderly(t *testing.T) {
	tests := []struct {
		name    string
		account string
		token   string
		params  string
	}{
		{
			name:    "simulate failed for tenderly",
			account: os.Getenv("account"),
			token:   os.Getenv("token"),
			params:  `{"network_id":"1","block_number":14365389,"transaction_index":0,"from":"0x7da5eacc8628f22d5e56ed0018751a8921942e38","input":"0xf8ca0f85174876e80083030d40947cad06b811b5d9d3ff197c1a046abcbc0efbcbc980b864d675fd260000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c626a0189c5d42fea496685df9488e8c2b761a278a7a6744383a9a6d7927dc11f70c29a02a948c7f03f80a30cf7fe5262c0f95068fa2c111b655d6a73cc223b151b6c0dc","to":"0x7cad06b811b5d9d3ff197c1a046abcbc0efbcbc9","gas":200000,"gas_price":"100000000000","value":"0","access_list":[],"generate_access_list":true,"save":false,"source":"dashboard"}`,
		},
		{
			name:    "simulate success for tenderly",
			account: os.Getenv("account"),
			token:   os.Getenv("token"),
			params:  `{"network_id":"1","block_number":14365440,"transaction_index":0,"from":"0x7da5eacc8628f22d5e56ed0018751a8921942e38","input":"0xd675fd260000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000c7","to":"0x7cad06b811b5d9d3ff197c1a046abcbc0efbcbc9","gas":200000,"gas_price":"100000000000","value":"0","access_list":[],"generate_access_list":true,"save":false,"source":"dashboard"}`,
		},
		{
			name:    "params error",
			account: "",
			token:   "",
			params:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulateRes, err := NewSimulate(tt.account, tt.token).SimulateGetter(SimulateTenderly).Simulate(tt.params)
			if err != nil {
				t.Fatalf("simulate failed for tenderly")
			}
			fmt.Println("simulateRes:", simulateRes)

			res := &providers.TenderlySimulateRes{}
			if err = json.Unmarshal([]byte(simulateRes), res); err != nil {
				t.Fatalf("unMarshal error: %s", err)
			}

			t.Log(res)
		})
	}
}
