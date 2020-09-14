package core

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/linger1216/go-front/echo-service/svc"
	"github.com/linger1216/go-front/geom"
	"github.com/linger1216/go-front/utils"
	"strings"
)

const (
	HeadCountKey = "X-Total-Count"
)

type EchoDBService struct {
	db  *sqlx.DB
	ddl *EchoDDL
}

func NewEchoDBService(db *sqlx.DB) svc.EchoServer {
	server := &EchoDBService{db, NewEchoDDL()}
	query := server.ddl.CreateTableDDL()
	if _, err := server.db.Exec(query); err != nil {
		panic(err)
	}
	query = server.ddl.IndexTableDDL()
	if _, err := server.db.Exec(query); err != nil {
		panic(err)
	}
	return server
}

func (f *EchoDBService) CreateEcho(ctx context.Context, in *svc.CreateEchoRequest) (*svc.CreateEchoResponse, error) {
	query, args := f.ddl.Upsert(in.Echos...)
	fmt.Printf("%s\n", query)

	_, err := f.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	for i := range in.Echos {
		ids = append(ids, in.Echos[i].Id)
	}
	return &svc.CreateEchoResponse{Ids: ids}, nil
}

func (f *EchoDBService) DeleteEcho(ctx context.Context, in *svc.DeleteEchoRequest) (*svc.DeleteEchoResponse, error) {
	query, args := f.ddl.Delete(in.Ids...)
	fmt.Printf("%s\n", query)
	_, err := f.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &svc.DeleteEchoResponse{}, nil
}

func (f *EchoDBService) UpdateEcho(ctx context.Context, in *svc.UpdateEchoRequest) (*svc.UpdateEchoResponse, error) {
	query, args := f.ddl.Upsert(in.Echos...)
	fmt.Printf("%s\n", query)
	_, err := f.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &svc.UpdateEchoResponse{}, nil
}

func (f *EchoDBService) ListEcho(ctx context.Context, in *svc.ListEchoRequest) (*svc.ListEchoResponse, error) {
	resp := &svc.ListEchoResponse{}
	query, args := f.ddl.List(in)
	fmt.Printf("%s\n", query)
	if in.Header {
		count := int64(0)
		err := f.db.Get(&count, query, args...)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, svc.ErrNotFound
		}
		resp.Headers = append(resp.Headers, &svc.KVResponse{
			Key:   HeadCountKey,
			Value: utils.Int64ToString(count),
		})
	} else {
		ret, err := f.query(query, args...)
		if err != nil {
			return nil, err
		}
		resp.Echos = ret
	}

	if len(resp.Echos) == 0 {
		return nil, svc.ErrNotFound
	}
	return resp, nil
}

func (f *EchoDBService) query(query string, args ...interface{}) ([]*svc.Echo, error) {
	rows, err := f.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	ret := make([]*svc.Echo, 0)
	for rows.Next() {
		line := make(map[string]interface{})
		err = rows.MapScan(line)
		if err != nil {
			return nil, err
		}
		if tc, err := transEcho("", line); err == nil && tc != nil {
			ret = append(ret, tc)
		}
	}

	if len(ret) == 0 {
		return nil, svc.ErrNotFound
	}
	return ret, nil
}

func transEcho(prefix string, m map[string]interface{}) (*svc.Echo, error) {
	ret := &svc.Echo{}

	if v, ok := m[prefix+"id"]; ok {
		ret.Id = utils.ToString(v)
	}

	if v, ok := m[prefix+"age"]; ok {
		ret.Age = int(utils.ToInt64(v))
	}

	if v, ok := m[prefix+"name"]; ok {
		ret.Name = utils.ToString(v)
	}

	if v, ok := m[prefix+"geometry"]; ok {
		if buf := utils.ToString(v); len(buf) > 0 {
			geometry := &geom.Geometry{}
			err := jsoniter.ConfigFastest.Unmarshal([]byte(buf), geometry)
			if err != nil {
				return nil, err
			}
			ret.Geometry = geometry
		}
	}

	if v, ok := m[prefix+"books"]; ok {
		ret.Books = strings.Split(utils.ToString(v), ",")
	}

	if v, ok := m[prefix+"tags"]; ok {
		arr := strings.Split(utils.ToString(v), ",")
		ret.Tags = make([]int, 0, len(arr))
		for i := range arr {
			ret.Tags = append(ret.Tags, int(utils.StringToInt64(arr[i])))
		}
	}

	if v, ok := m[prefix+"create_time"]; ok {
		ret.CreateTime = utils.ToInt64(v)
	}
	if v, ok := m[prefix+"update_time"]; ok {
		ret.UpdateTime = utils.ToInt64(v)
	}

	return ret, nil
}

func (f *EchoDBService) GetEcho(ctx context.Context, in *svc.GetEchoRequest) (*svc.GetEchoResponse, error) {
	query, args := f.ddl.Get(in.Ids...)
	fmt.Printf("%s\n", query)
	ret, err := f.query(query, args...)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, svc.ErrNotFound
	}
	return &svc.GetEchoResponse{Echos: ret}, nil
}

func (f *EchoDBService) Close() error {
	return nil
}
