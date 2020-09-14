package core

import (
	"context"
	"github.com/linger1216/go-front/echo-service/svc"
)

type EchoCacheService struct {
}

func NewEchoCacheService() svc.EchoServer {
	return &EchoCacheService{}
}

func (f *EchoCacheService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	return nil, nil
}

func (f *EchoCacheService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	return nil, nil
}

func (f *EchoCacheService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	return nil, nil
}

func (f *EchoCacheService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	return nil, nil
}

func (f *EchoCacheService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	return nil, nil
}

func (f *EchoCacheService) Close() error {
	return nil
}
