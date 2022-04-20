package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

type Tenderly struct {
	URL     string
	Project string
	Token   string
	Limiter *rate.Limiter
}

func NewTenderly(account, project, token string, tps int) *Tenderly {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second*1), tps)
	//url := fmt.Sprintf("https://api.tenderly.co/api/v1/account/%s/project/%s", account, project)
	url := fmt.Sprintf("https://api.tenderly.co/api/v1/account/%s/project", account)
	return &Tenderly{URL: url, Project: project, Token: token, Limiter: rateLimiter}
}

func (t *Tenderly) ReqJsonRpc(rpc string, params *RpcParams) (string, error) {
	_ = t.Limiter.Wait(context.Background())
	//https://rpc.tenderly.co/fork/forkId

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	bEnv, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return "", err
	}

	res, err := NewNet(rpc, header, mapParam).Request(PostTy)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (t *Tenderly) AddProject(name string) (string, error) {
	// https://api.tenderly.co/api/v1/account/zck/project
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	project := &struct {
		Name string
	}{
		name,
	}

	bEnv, err := json.Marshal(project)
	if err != nil {
		return "", err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return "", err
	}

	url := t.URL
	res, err := NewNet(url, header, mapParam).Request(PostTy)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (t *Tenderly) RenameProject(name string) (string, error) {
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	project := &struct {
		Name string
	}{
		name,
	}

	bEnv, err := json.Marshal(project)
	if err != nil {
		return "", err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return "", err
	}

	url := t.URL + "/" + t.Project
	res, err := NewNet(url, header, mapParam).Request(PostTy)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (t *Tenderly) AddForkEnv(chainId, name string) (string, error) {
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
		return "", err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return "", err
	}

	url := t.URL + "/" + t.Project + "/fork"
	res, err := NewNet(url, header, mapParam).Request(PostTy)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (t *Tenderly) RenameForkEnv(forkId, chainId, name string) (string, error) {
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
		return "", err
	}

	mapParam := make(map[string]interface{})
	err = json.Unmarshal(bEnv, &mapParam)
	if err != nil {
		return "", err
	}

	url := t.URL + "/" + t.Project + "/fork/" + forkId
	res, err := NewNet(url, header, mapParam).Request(PutTy)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (t *Tenderly) DeleteForkEnv(id string) (string, error) {
	//  https: //api.tenderly.co/api/v1/account/zck/project/project/fork/30c50584-b7a2-4819-998a-e9ef85749575
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	url := t.URL + "/" + t.Project + "/fork/" + id
	res, err := NewNet(url, header, nil).Request(DeleteTy)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (t *Tenderly) SimulateTxForFork(forkId, params string) (string, error) {
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

	url := t.URL + "/" + t.Project + "/fork/" + forkId + "/simulate"
	return NewNet(url, header, mapParam).Request(PostTy)
}

func (t *Tenderly) Simulate(params string) (string, error) {
	_ = t.Limiter.Wait(context.Background())
	header := make(map[string]string)
	header["X-Access-Key"] = t.Token
	header["Content-Type"] = "application/json"

	mapParam := make(map[string]interface{})
	err := json.Unmarshal([]byte(params), &mapParam)
	if err != nil {
		return "", err
	}

	url := t.URL + "/" + t.Project + "/simulate"
	return NewNet(url, header, mapParam).Request(PostTy)
}
