package meta

import (
	"github.com/linger1216/go-front/utils"
	"testing"
)

func TestTableInfo_toGolangDefinition(t *testing.T) {
	pg := NewPostgresMeta(utils.NewPostgres(&utils.PostgresConfig{
		Uri:     "postgres://lid.guan:@localhost:15432/zhigan?sslmode=disable",
		MaxIdle: 0,
		MaxOpen: 0,
	}))

	tables, _ := pg.GetInfo()
	for _, v := range tables {
		v.toGolangDefinition(true)
	}
}
