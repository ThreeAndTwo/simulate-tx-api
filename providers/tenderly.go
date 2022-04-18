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
	Project string
	Token   string
	Limiter *rate.Limiter
}

func NewTenderSimulate(account, project, token string, tps int) *TenderSimulate {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second*1), tps)
	url := fmt.Sprintf("https://api.tenderly.co/api/v1/account/%s/project/%s", account, project)
	return &TenderSimulate{URL: url, Project: project, Token: token, Limiter: rateLimiter}
}

func (t *TenderSimulate) AddForkEnv(chainId, name string) error {
	// https://api.tenderly.co/api/v1/account/zck/project/project/fork
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	env := &AddForEnv{
		NetworkId: chainId,
		Alias:     name,
	}

	bEnv, err := json.Marshal(env)
	if err != nil {
		return err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return err
	}

	_, err = NewNet(t.URL, header, mapParam).Request(PostTy)
	if err != nil {
		return err
	}
	return nil
}

func (t *TenderSimulate) RenameForkEnv(forkId, chainId, name string) error {
	//https://api.tenderly.co/api/v1/account/zck/project/project/fork/8a5d135f-447b-4fa3-9ed2-bcb91498d39a
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	env := &AddForEnv{
		NetworkId: chainId,
		Alias:     name,
	}

	bEnv, err := json.Marshal(env)
	if err != nil {
		return err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return err
	}

	url := t.URL + "/fork/" + forkId
	_, err = NewNet(url, header, mapParam).Request(PutTy)
	if err != nil {
		return err
	}
	return nil
}

func (t *TenderSimulate) DeleteForkEnv(id string) error {
	//  https: //api.tenderly.co/api/v1/account/zck/project/project/fork/30c50584-b7a2-4819-998a-e9ef85749575
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	url := t.URL + "/fork/" + id
	_, err := NewNet(url, header, nil).Request(DeleteTy)
	if err != nil {
		return err
	}
	return nil
}

func (t *TenderSimulate) SimulateTxForFork(forkId, params string) (string, error) {
	//https: //api.tenderly.co/api/v1/account/zck/project/project/fork/8a5d135f-447b-4fa3-9ed2-bcb91498d39a/simulate
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	mapParam := make(map[string]interface{})
	err := json.Unmarshal([]byte(params), &mapParam)
	if err != nil {
		return "", err
	}

	url := t.URL + "/fork/" + forkId + "/simulate"
	return NewNet(url, header, mapParam).Request(PostTy)
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

	url := t.URL + "/simulate"
	return NewNet(url, header, mapParam).Request(PostTy)
}
