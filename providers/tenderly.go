package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

type TenderSimulate struct {
	URL     string
	Token   string
	Limiter *rate.Limiter
}

func NewTenderSimulate(account string, token string, tps int) *TenderSimulate {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second*1), tps)
	url := fmt.Sprintf("https://api.tenderly.co/api/v1/account/%s/project/mywallet/simulate", account)
	return &TenderSimulate{URL: url, Token: token, Limiter: rateLimiter}
}

func (t *TenderSimulate) Simulate(params string) (string, error) {
	_ = t.Limiter.Wait(context.Background())
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
