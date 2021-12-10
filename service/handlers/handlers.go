package handlers

import (
	"context"
	"fmt"
	"github.com/linger1216/go-front/core"
	"github.com/linger1216/go-front/service/svc"
	"github.com/linger1216/go-front/utils"
	"github.com/linger1216/go-front/utils/config"
	coordTransform "github.com/qichengzx/coordtransform"
	"strings"
	"time"
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
		cacheCfg.UseCache = false
	}

	ddlCfg := &core.DDLConfig{}
	if err := reader.ScanKey("ddl", ddlCfg); err != nil {
		panic(err)
	}

	return &EchoService{db: core.NewEchoDBService(pg, ddlCfg), cache: core.NewEchoCacheService(cacheCfg)}
}

func (f *EchoService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	if in == nil || len(in.Echos) == 0 {
		return nil, svc.ErrInvalidPara
	}
	return f.db.CreateEcho(ctx, in)
}

func (f *EchoService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	if in == nil || len(in.Ids) == 0 {
		return nil, svc.ErrInvalidPara
	}
	return f.db.DeleteEcho(ctx, in)
}

func (f *EchoService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	if in == nil || len(in.Echos) == 0 {
		return nil, svc.ErrInvalidPara
	}
	resp, err := f.db.UpdateEcho(ctx, in)
	if err != nil {
		return nil, err
	}
	_ = f.cache.UpdateEcho(ctx, in)
	return resp, nil
}

func (f *EchoService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	if in == nil {
		return nil, svc.ErrInvalidPara
	}

	if in.Point != nil {
		if len(in.Point.Coordinates) < 2 || !utils.LnglatValid(in.Point.Coordinates[0], in.Point.Coordinates[1]) {
			return nil, svc.ErrInvalidPara
		}

		if len(in.Point.SpatialReference) == 0 {
			in.Point.SpatialReference = "gcj02"
		}

		switch strings.ToLower(in.Point.SpatialReference) {
		case "wgs84":
			in.Point.Coordinates[0], in.Point.Coordinates[1] =
				coordTransform.WGS84toGCJ02(in.Point.Coordinates[0], in.Point.Coordinates[1])
		case "bd09":
			in.Point.Coordinates[0], in.Point.Coordinates[1] =
				coordTransform.BD09toGCJ02(in.Point.Coordinates[0], in.Point.Coordinates[1])
		case "gcj02":
		}
	}

	if in.EndTime == 0 {
		in.EndTime = time.Now().Unix()
	}

	if in.PageSize == 0 {
		in.PageSize = 10
	}

	if in.Radius == 0 {
		in.Radius = 1000
	}

	if resp, err := f.cache.ListEcho(ctx, in); err == nil && resp != nil {
		fmt.Printf("ListEcho in cache\n")
		return resp, nil
	}
	resp, err := f.db.ListEcho(ctx, in)
	if err != nil {
		return nil, err
	}
	_ = f.cache.UpdateListEcho(ctx, in, resp)
	return resp, nil
}

func (f *EchoService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	if in == nil || len(in.Ids) == 0 {
		return nil, svc.ErrInvalidPara
	}
	if resp, err := f.cache.GetEcho(ctx, in); err == nil && resp != nil {
		fmt.Printf("GetEcho in cache\n")
		return resp, nil
	}
	resp, err := f.db.GetEcho(ctx, in)
	if err != nil {
		return nil, err
	}
	_ = f.cache.UpdateGetEcho(ctx, in, resp)
	return resp, nil
}

func (f *EchoService) Close() error {
	_ = f.cache.Close()
	return f.db.Close()
}
