package handlers

import (
	"context"
	"github.com/linger1216/go-front/core"
	"github.com/linger1216/go-front/echo-service/svc"
)

type EchoService struct {
	proxy *core.EchoImplService
}

func NewService() svc.EchoServer {
	return &EchoService{}
}

func (f *EchoService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	return f.proxy.CreateEcho(ctx, in)
}

func (f *EchoService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	return f.proxy.DeleteEcho(ctx, in)
}

func (f *EchoService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	return f.proxy.UpdateEcho(ctx, in)
}

func (f *EchoService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	return f.proxy.ListEcho(ctx, in)
}

func (f *EchoService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	return f.proxy.GetEcho(ctx, in)
}

func (f *EchoService) Close() error {
	return f.proxy.Close()
}
