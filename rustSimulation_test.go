package simulate_tx_api

import "testing"

const host = "http://127.0.0.1:10000"

func TestNewSimulate_simulate(t *testing.T) {
	tests := []struct {
		name  string
		param string
	}{
		{
			name:  "simulation",
			param: "{\n  \"chainId\": 1,\n  \"from\": \"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045\",\n  \"to\": \"0x66fc62c1748e45435b06cf8dd105b73e9855f93e\",\n  \"data\": \"0xffa2ca3b44eea7c8e659973cbdf476546e9e6adfd1c580700537e52ba7124933a97904ea000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001d0e30db00300ffffffffffffc02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000186a0\",\n  \"gasLimit\": 500000,\n  \"value\": \"100000\",\n  \"blockNumber\": 16784600\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulations, err := NewSimulate(host, "", "", "", 5).SimulateGetter(SimulateCoinSummer).Simulations(tt.param)
			if err != nil {
				t.Errorf("simulation error: %s", err)
				return
			}
			t.Log(simulations)
		})
	}
}

func TestNewSimulate_simulateBundle(t *testing.T) {
	tests := []struct {
		name  string
		param string
	}{
		{
			name:  "simulate-bundle",
			param: "[\n  {\n    \"chainId\": 1,\n    \"from\": \"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045\",\n    \"to\": \"0x66fc62c1748e45435b06cf8dd105b73e9855f93e\",\n    \"data\": \"0xffa2ca3b44eea7c8e659973cbdf476546e9e6adfd1c580700537e52ba7124933a97904ea000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001d0e30db00300ffffffffffffc02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000186a0\",\n    \"gasLimit\": 500000,\n    \"value\": \"100000\",\n    \"blockNumber\": 16784600\n  }\n]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulations, err := NewSimulate(host, "", "", "", 5).SimulateGetter(SimulateCoinSummer).BundledSimulations(tt.param)
			if err != nil {
				t.Errorf("simulation error: %s", err)
				return
			}
			t.Log(simulations)
		})
	}
}
