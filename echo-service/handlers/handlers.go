package handlers

import (
	"context"
	"github.com/linger1216/go-front/core"
	"github.com/linger1216/go-front/echo-service/svc"
	"github.com/linger1216/go-front/utils"
	"github.com/linger1216/go-front/utils/config"
)

type EchoService struct {
	db    svc.EchoServer
	local svc.EchoServer
}

func NewService(reader config.Reader) svc.EchoServer {
	cfg := &utils.PostgresConfig{}
	if err := reader.ScanKey("postgres", cfg); err != nil {
		panic(err)
	}
	pg := utils.NewPostgres(cfg)
	return &EchoService{db: core.NewEchoDBService(pg)}
}

func (f *EchoService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	return f.db.CreateEcho(ctx, in)
}

func (f *EchoService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	return f.db.DeleteEcho(ctx, in)
}

func (f *EchoService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	return f.db.UpdateEcho(ctx, in)
}

func (f *EchoService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	if in.PageSize == 0 {
		in.PageSize = 10
	}
	return f.db.ListEcho(ctx, in)
}

func (f *EchoService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	return f.db.GetEcho(ctx, in)
}

func (f *EchoService) Close() error {
	return f.db.Close()
}
