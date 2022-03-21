package providers

import (
	"encoding/json"
	"fmt"
)

type TenderSimulate struct {
	URL   string
	Token string
}

func NewTenderSimulate(account string, token string) *TenderSimulate {
	url := fmt.Sprintf("https://api.tenderly.co/api/v1/account/%s/project/mywallet/simulate", account)
	return &TenderSimulate{URL: url, Token: token}
}

func (t *TenderSimulate) Simulate(params string) (string, error) {
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	mapParam := make(map[string]interface{})
	err := json.Unmarshal([]byte(params), &mapParam)
	if err != nil {
		return "", err
	}

	return NewNet(t.URL, header, mapParam).Request()
}
