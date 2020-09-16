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
	cache svc.EchoCacheServer
}

func NewService(reader config.Reader) svc.EchoServer {
	dbCfg := &utils.PostgresConfig{}
	if err := reader.ScanKey("postgres", dbCfg); err != nil {
		panic(err)
	}
	pg := utils.NewPostgres(dbCfg)

	cacheCfg := &core.EchoCacheConfig{}
	if err := reader.ScanKey("cache", cacheCfg); err != nil {
		panic(err)
	}
	return &EchoService{db: core.NewEchoDBService(pg), cache: core.NewEchoCacheService(cacheCfg)}
}

func (f *EchoService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	return f.db.CreateEcho(ctx, in)
}

func (f *EchoService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	return f.db.DeleteEcho(ctx, in)
}

func (f *EchoService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	resp, err := f.db.UpdateEcho(ctx, in)
	if err != nil {
		return nil, err
	}
	_ = f.cache.UpdateEcho(ctx, in)
	return resp, nil
}

func (f *EchoService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	if in.PageSize == 0 {
		in.PageSize = 10
	}

	if in.Radius == 0 {
		in.Radius = 1000
	}

	if resp, err := f.cache.ListEcho(ctx, in); err == nil && resp != nil {
		return resp, nil
	}
	resp, err := f.db.ListEcho(ctx, in)
	if err != nil {
		_ = f.cache.UpdateListEcho(ctx, in, resp)
	}
	return resp, nil
}

func (f *EchoService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	if resp, err := f.cache.GetEcho(ctx, in); err == nil && resp != nil {
		return resp, nil
	}
	resp, err := f.db.GetEcho(ctx, in)
	if err != nil {
		_ = f.cache.UpdateGetEcho(ctx, in, resp)
	}
	return resp, nil
}

func (f *EchoService) Close() error {
	_ = f.cache.Close()
	return f.db.Close()
}
