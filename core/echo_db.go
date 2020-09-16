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

func NewEchoDDL() *EchoDDL {
	ret := &EchoDDL{Name: "echo_test"}
	ret.columns = append(ret.columns, &MetaColumn{Name: "id", Type: "character varying", Primary: true})
	ret.columns = append(ret.columns, &MetaColumn{Name: "age", Type: "integer"})
	ret.columns = append(ret.columns, &MetaColumn{Name: "name", Type: "character varying"})
	ret.columns = append(ret.columns, &MetaColumn{Name: "geometry", Type: "geometry(Geometry,4326)", Index: true})
	ret.columns = append(ret.columns, &MetaColumn{Name: "books", Type: "character varying[]"})
	ret.columns = append(ret.columns, &MetaColumn{Name: "tags", Type: "integer[]"})
	ret.columns = append(ret.columns, &MetaColumn{Name: "create_time", Type: "bigint", Default: `(date_part('epoch'::text, now()))::bigint`})
	ret.columns = append(ret.columns, &MetaColumn{Name: "update_time", Type: "bigint", Default: `(date_part('epoch'::text, now()))::bigint`})
	return ret
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

type MetaColumn struct {
	Name    string `json:"Name"`
	Type    string `json:"type"` // character varying, bigint, integer, geometry(Geometry,4326), character varying[], integer[]
	Primary bool   `json:"primary"`
	Index   bool   `json:"index"`
	Unique  bool   `json:"unique"`
	Default string `json:"default"`
}

func (m *MetaColumn) ColumnDDL() string {
	var primary string
	if m.Primary {
		primary = "primary key"
	}

	var defaultVal string
	if len(m.Default) > 0 {
		defaultVal = "default " + m.Default
	}

	return fmt.Sprintf("%s %s %s %s", m.Name, m.Type, primary, defaultVal)
}

func (m *MetaColumn) Select() string {
	switch m.Type {
	case `character varying`:
		return m.Name
	case `bigint`:
		return m.Name
	case `integer`:
		return m.Name
	case `geometry(Geometry,4326)`:
		return fmt.Sprintf("st_asgeojson(%s) as %s", m.Name, m.Name)
	case `character varying[]`:
		return fmt.Sprintf("array_to_string(%s, ',') as %s", m.Name, m.Name)
	case `integer[]`:
		return fmt.Sprintf("array_to_string(%s, ',') as %s", m.Name, m.Name)
	}
	return ""
}

func (m *MetaColumn) IndexDDL(table string) string {
	if m.Primary {
		return ""
	}
	unique := ""
	if m.Unique {
		unique = "unique"
	}
	engine := ""
	switch m.Type {
	case "character varying", "bigint", "integer", "character varying[]", "integer[]":
		engine = fmt.Sprintf("btree(%s)", m.Name)
	case "geometry(Geometry,4326)":
		engine = fmt.Sprintf("gist (geography(%s))", m.Name)
	}
	return fmt.Sprintf("create %s index if not exists %s_%s_index ON %s using %s;", unique, table, m.Name, table, engine)
}

type EchoDDL struct {
	Name    string
	columns []*MetaColumn
}

func (m *EchoDDL) Select() string {
	arr := make([]string, len(m.columns))
	for i := range m.columns {
		arr[i] = m.columns[i].Select()
	}
	return strings.Join(arr, ",")
}

func (m *EchoDDL) Table() string {
	return m.Name
}

func (m *EchoDDL) ColumnsString() string {
	arr := make([]string, len(m.columns))
	for i := range m.columns {
		arr[i] = m.columns[i].Name
	}
	return strings.Join(arr, ",")
}

func (m *EchoDDL) CreateTableDDL() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("create table if not exists %s", m.Name))
	buf.WriteString("(\n")
	for i := range m.columns {
		buf.WriteString(m.columns[i].ColumnDDL())
		if i < len(m.columns)-1 {
			buf.WriteByte(',')
		}
		buf.WriteByte('\n')
	}
	buf.WriteString(");\n")
	return buf.String()
}

func (m *EchoDDL) IndexTableDDL() []string {
	arr := make([]string, 0)
	for _, v := range m.columns {
		if v.Index {
			arr = append(arr, v.IndexDDL(m.Name))
		}
	}
	return arr
}

func (m *EchoDDL) DBPrimaryColumn() *MetaColumn {
	for i := range m.columns {
		if m.columns[i].Primary {
			return m.columns[i]
		}
	}
	return nil
}

func (m *EchoDDL) OnConflictDDL() string {
	primaryColumn := m.DBPrimaryColumn()
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("on conflict (%s)\n", primaryColumn.Name))
	buf.WriteString("do update set\n")

	for i, v := range m.columns {
		if v.Primary || v.Name == "create_time" {
			continue
		}
		if v.Name == "update_time" {
			buf.WriteString(fmt.Sprintf("update_time = GREATEST(%s.update_time, excluded.update_time)", m.Name))
		} else {
			buf.WriteString(fmt.Sprintf("%s = excluded.%s", v.Name, v.Name))
		}
		if i < len(m.columns)-1 {
			buf.WriteString(",")
		} else {
			buf.WriteString(";")
		}
		buf.WriteString("\n")
	}
	return buf.String()

}
