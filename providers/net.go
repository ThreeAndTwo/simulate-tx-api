package providers

import (
	"encoding/json"
	"github.com/deng00/req"
	"strings"
)

type Net struct {
	Url    string
	Header map[string]string
	Params map[string]interface{}
	IsJson bool
}

type NetType string

const (
	GetTy    NetType = "get"
	PostTy   NetType = "post"
	DeleteTy NetType = "delete"
	PutTy    NetType = "put"
)

func NewNet(url string, header map[string]string, params map[string]interface{}) *Net {
	return &Net{Url: url, Header: header, Params: params}
}

func (n *Net) Request(netType NetType) (string, error) {
	reqHeader, hasJson := n.initHeader()
	reqParams := n.initParam()

	if hasJson {
		n.IsJson = hasJson
	}

	switch netType {
	case GetTy:
		return n.get(reqHeader)
	case PostTy:
		return n.post(reqHeader, reqParams)
	case DeleteTy:
		return n.delete(reqHeader)
	case PutTy:
		return n.put(reqHeader, reqParams)
	default:
		return n.get(reqHeader)
	}
}

func (n *Net) initParam() req.Param {
	reqParams := req.Param{}
	for k, v := range n.Params {
		reqParams[k] = v
	}
	return reqParams
}

func (n *Net) get(header req.Header) (string, error) {
	resp, err := req.Get(n.Url, header)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (n *Net) post(header req.Header, param req.Param) (string, error) {
	var reqResp = &req.Resp{}
	var err error

	if n.IsJson {
		jsonParam, _ := json.Marshal(param)
		reqResp, err = req.Post(n.Url, header, jsonParam)
	} else {
		reqResp, err = req.Post(n.Url, header, param)
	}
	return reqResp.String(), err
}

func (n *Net) delete(header req.Header) (string, error) {
	reqResp, err := req.Delete(n.Url, header)
	return reqResp.String(), err
}

func (n *Net) put(header req.Header, param req.Param) (string, error) {
	var reqResp = &req.Resp{}
	var err error

	if n.IsJson {
		jsonParam, _ := json.Marshal(param)
		reqResp, err = req.Put(n.Url, header, jsonParam)
	} else {
		reqResp, err = req.Put(n.Url, header, param)
	}
	return reqResp.String(), err
}

func (n *Net) initHeader() (req.Header, bool) {
	authHeader := req.Header{}
	hasJson := false

	for k, v := range n.Header {
		authHeader[k] = v
		if hasJsonInHeader(k, v) {
			hasJson = true
		}
	}
	return authHeader, hasJson
}

func hasJsonInHeader(key, value string) bool {
	return strings.ToLower(key) == "content-type" && strings.Contains(strings.ToLower(value), "json")
}
