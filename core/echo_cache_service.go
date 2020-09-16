package core

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/linger1216/go-front/echo-service/svc"
	"github.com/linger1216/go-front/utils"
)

type EchoCacheConfig struct {
	UseCache    bool `yaml:"useCache"`
	CacheExpire int  `yaml:"cacheExpire"`
	CacheSize   int  `yaml:"cacheSize"`
}

type EchoCacheService struct {
	enable bool
	local  *utils.Ristretto
}

func NewEchoCacheService(cfg *EchoCacheConfig) svc.EchoCacheServer {
	ret := &EchoCacheService{enable: cfg.UseCache}
	if cfg.UseCache {
		size := cfg.CacheSize * (1 << 20)
		ret.local = utils.NewRistretto(int64(size))
	}
	return ret
}

func (e *EchoCacheService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) error {
	return nil
}

func genListEchoRequestKey(in *svc.ListEchoRequest) (string, error) {
	buf, err := jsoniter.ConfigFastest.Marshal(in)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("list:%s", string(buf)), nil
}

func (e *EchoCacheService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	if !e.enable {
		return nil, nil
	}
	key, err := genListEchoRequestKey(in)
	if err != nil {
		return nil, err
	}

	if len(key) == 0 {
		return nil, nil
	}

	resp, err := e.local.Get(key)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return resp.(*svc.ListEchoResponse), nil
}

func (e *EchoCacheService) UpdateListEcho(ctx context.Context, in *svc.ListEchoRequest, resp *svc.ListEchoResponse) error {
	if !e.enable {
		return nil
	}
	key, err := genListEchoRequestKey(in)
	if err != nil {
		return err
	}
	return e.local.Set(key, resp)
}

func genGetEchoRequestKey(in *svc.GetEchoRequest) (string, error) {
	buf, err := jsoniter.ConfigFastest.Marshal(in)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("get:%s", string(buf)), nil
}

func (e *EchoCacheService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	if !e.enable {
		return nil, nil
	}
	key, err := genGetEchoRequestKey(in)
	if err != nil {
		return nil, err
	}

	if len(key) == 0 {
		return nil, nil
	}

	resp, err := e.local.Get(key)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}
	return resp.(*svc.GetEchoResponse), nil
}

func (e *EchoCacheService) UpdateGetEcho(ctx context.Context, in *svc.GetEchoRequest, resp *svc.GetEchoResponse) error {
	if !e.enable {
		return nil
	}
	key, err := genGetEchoRequestKey(in)
	if err != nil {
		return err
	}
	return e.local.Set(key, resp)
}

func (e *EchoCacheService) Close() error {
	return nil
}
