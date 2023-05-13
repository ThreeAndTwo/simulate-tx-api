package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type PrivateTxSimulator struct {
	Host    string
	Limiter *rate.Limiter
}

func (p *PrivateTxSimulator) AddProject(name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) RenameProject(name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) AddForkEnv(chainId, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) RenameForkEnv(forkId, chainId, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) DeleteForkEnv(forkId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) SimulateTxForFork(forkId, params string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) Simulations(params string) (string, error) {
	_ = p.Limiter.Wait(context.Background())
	header := p.getReqHeader()

	mapParam := make(map[string]interface{})
	err := json.Unmarshal([]byte(params), &mapParam)
	if err != nil {
		return "", err
	}

	url := p.Host + "/simulate"
	return NewNet(url, header, mapParam).Request(PostTy)
}

func (p *PrivateTxSimulator) BundledSimulations(params string) (string, error) {
	_ = p.Limiter.Wait(context.Background())
	url := p.Host + "/simulate-bundle"
	return reqPost(url, params)
}

func reqPost(url, data string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("req error:%s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parse resp data error: %s", err)
	}
	return string(body), nil
}

func (p *PrivateTxSimulator) ReqJsonRpc(rpc string, params *RpcParams) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PrivateTxSimulator) getReqHeader() (header map[string]string) {
	header = make(map[string]string)
	header["Content-Type"] = "application/json"
	return
}

func NewPrivateTxSimulator(host string, tps int) *PrivateTxSimulator {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second*1), tps)
	host += "/api/v1"
	return &PrivateTxSimulator{Host: host, Limiter: rateLimiter}
}
