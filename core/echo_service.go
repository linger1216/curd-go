package core

import (
	"context"
	"github.com/linger1216/go-front/echo-service/svc"
)

type EchoImplService struct {
}

func NewService() svc.EchoServer {
	return &EchoImplService{}
}

func (f *EchoImplService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	return nil, nil
}

func (f *EchoImplService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	return nil, nil
}

func (f *EchoImplService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	return nil, nil
}

func (f *EchoImplService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	return nil, nil
}

func (f *EchoImplService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	return nil, nil
}

func (f *EchoImplService) Close() error {
	return nil
}
