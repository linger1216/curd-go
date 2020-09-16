package core

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/lib/pq"
	"github.com/linger1216/go-front/echo-service/svc"
	"github.com/linger1216/go-front/utils"
	"strings"
	"time"
)

type EchoDDL struct {
}

func NewEchoDDL() *EchoDDL {
	ret := &EchoDDL{}
	return ret
}

func (m *EchoDDL) Select() string {
	return `id, name, age, st_asgeojson(geometry) as geometry, array_to_string(tags, ',') as tags, array_to_string(books, ',') as books, create_time, update_time`
}

func (m *EchoDDL) Table() string {
	return `echo_table`
}

func (m *EchoDDL) ColumnsString() string {
	return `id,age,name,geometry,books,tags,create_time,update_time`
}

func (m *EchoDDL) CreateTableDDL() string {
	return fmt.Sprintf(`
create table if not exists %s
(
	id       character varying primary key,
	age      integer,
	name     character varying,
  geometry geometry(Geometry,4326),
  books    character varying[],
  tags     integer[],
  create_time   bigint default (date_part('epoch'::text, now()))::bigint,
  update_time   bigint default (date_part('epoch'::text, now()))::bigint
);`, m.Table())
}

func (m *EchoDDL) IndexTableDDL() string {
	// 需要加上if not exists
	return ` CREATE UNIQUE INDEX if not exists echo_table_pkey ON public.echo_table USING btree (id);`
}

func (m *EchoDDL) OnConflictDDL() string {
	return fmt.Sprintf(`
on conflict (id) 
do update set 
age = excluded.age, 
name = excluded.name, 
geometry = excluded.geometry, 
books = excluded.books, 
tags = excluded.tags,
update_time = GREATEST(%s.update_time, excluded.update_time);
`, m.Table())
}

func (m *EchoDDL) Upsert(echos ...*svc.Echo) (string, []interface{}) {
	cols := strings.Split(m.ColumnsString(), ",")
	size := len(echos)
	values := make([]string, 0, size)
	args := make([]interface{}, 0, size*len(cols))
	for i, v := range echos {
		if len(v.Id) == 0 {
			v.Id = utils.Generate()
		}

		var createTime, updateTime int64
		if v.CreateTime == 0 {
			createTime = time.Now().Unix()
		} else {
			createTime = v.CreateTime
		}

		if v.UpdateTime == 0 {
			updateTime = time.Now().Unix()
		} else {
			updateTime = v.UpdateTime
		}

		geometry, err := jsoniter.ConfigFastest.Marshal(v.Geometry)
		if err != nil {
			panic(err)
		}
		values = append(values, utils.ValuePlaceHolderAndGeometry(i*len(cols), len(cols), 4))
		args = append(args, v.Id, v.Age, v.Name, geometry, pq.Array(v.Books), pq.Array(v.Tags), createTime, updateTime)
	}

	query := fmt.Sprintf(`insert into %s (%s) values %s %s`, m.Table(), m.ColumnsString(),
		strings.Join(values, ","), m.OnConflictDDL())
	return query, args
}

func (m *EchoDDL) List(in *svc.ListEchoRequest) (string, []interface{}) {
	firstCond := true
	var buffer bytes.Buffer
	if in.Header {
		buffer.WriteString(fmt.Sprintf("select count(1) from %s", m.Table()))
	} else {
		buffer.WriteString(fmt.Sprintf("select %s from %s", m.Select(), m.Table()))
	}

	if len(in.Ages) > 0 {
		query := fmt.Sprintf("%s age in (%s)", utils.CondSql(firstCond), utils.SqlIntegerIn(in.Ages...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.Names) > 0 {
		query := fmt.Sprintf("%s name in (%s)", utils.CondSql(firstCond), utils.SqlStringIn(in.Names...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.Books) > 0 {
		query := fmt.Sprintf("%s books @> %s", utils.CondSql(firstCond), utils.SqlStringArray(in.Books...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.Tags) > 0 {
		query := fmt.Sprintf("%s tags @> %s", utils.CondSql(firstCond), utils.SqlIntegerArray(in.Tags...))
		buffer.WriteString(query)
		firstCond = false
	}

	if in.Point != nil && in.Radius > 0 {
		query := fmt.Sprintf("%s %s", utils.CondSql(firstCond),
			utils.SqlWithIn(in.Point.Coordinates[0], in.Point.Coordinates[1], int(in.Radius)))
		buffer.WriteString(query)
		firstCond = false
	}

	if in.StartTime > 0 && in.EndTime > 0 {
		query := fmt.Sprintf("%s update_time between '%d' and '%d' ", utils.CondSql(firstCond), in.StartTime, in.EndTime)
		buffer.WriteString(query)
		firstCond = false
	}

	if !in.Header {
		query := fmt.Sprintf(" offset %d limit %d", in.CurrentPage*in.PageSize, in.PageSize)
		buffer.WriteString(query)
	}

	buffer.WriteString(";")
	return buffer.String(), nil
}

func (m *EchoDDL) Delete(ids ...string) (string, []interface{}) {
	query := fmt.Sprintf("delete from %s where %s in (%s);", m.Table(), "id", utils.SqlStringIn(ids...))
	return query, nil
}

func (m *EchoDDL) Get(ids ...string) (string, []interface{}) {
	query := fmt.Sprintf("select %s from %s where %s in (%s);", m.Select(), m.Table(), "id", utils.SqlStringIn(ids...))
	return query, nil
}
