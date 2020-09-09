package handlers

import (
	"context"
	"github.com/linger1216/go-front/echo-service/svc"
)

type EchoService struct {
}

func NewService() svc.EchoServer {
	return &EchoService{}
}

func (f *EchoService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	return nil, nil
}

func (f *EchoService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	return nil, nil
}

func (f *EchoService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	return nil, nil
}

func (f *EchoService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	return nil, nil
}

func (f *EchoService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	return nil, nil
}

func (f *EchoService) Close() error {
	return nil
}
