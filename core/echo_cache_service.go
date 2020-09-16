package core

import (
	"context"
	"github.com/linger1216/go-front/echo-service/svc"
)

type EchoCacheConfig struct {
	UseCache    bool `yaml:"useCache"`
	CacheExpire int  `yaml:"cacheExpire"`
	CacheSize   int  `yaml:"cacheSize"`
}

type EchoCacheService struct {
}

func (e *EchoCacheService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) error {
	panic("implement me")
}

func (e *EchoCacheService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	panic("implement me")
}

func (e *EchoCacheService) UpdateListEcho(ctx context.Context, in *svc.ListEchoRequest, resp *svc.ListEchoResponse) error {
	panic("implement me")
}

func (e *EchoCacheService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	panic("implement me")
}

func (e *EchoCacheService) UpdateGetEcho(ctx context.Context, in *svc.GetEchoRequest, resp *svc.GetEchoResponse) error {
	panic("implement me")
}

func (e *EchoCacheService) Close() error {
	panic("implement me")
}

func NewEchoCacheService(cfg *EchoCacheConfig) svc.EchoCacheServer {
	return &EchoCacheService{}
}
