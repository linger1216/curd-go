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

	arr := make([]string, 0)
	tables, _ := pg.GetInfo()
	for _, v := range tables {
		if v.Name == "echo_table" {
			arr = append(arr, v.DBSelectColumn())
			arr = append(arr, v.DBColumns())
			arr = append(arr, v.DBCreateTableDDL(true))
			arr = append(arr, v.DBIndexTableDDL()...)
			arr = append(arr, v.DBOnConflictDDL())
			arr = append(arr, v.DBUpsert())
			arr = append(arr, v.DBList())
		}
	}
}
