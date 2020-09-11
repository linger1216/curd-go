package core

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/linger1216/go-front/echo-service/svc"
	"github.com/linger1216/go-front/utils"
	"strings"
	"time"
)

type MetaColumn struct {
	Name string `json:"Name"`
	// int64, int, string, geometry
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
	// 默认值
	Default string `json:"default"`
	// 索引语句
	Index  string `json:"index"`
	dbType string
}

func (m *MetaColumn) ColumnDDL() string {
	var primary string
	if m.Primary {
		primary = "primary key"
	}
	return fmt.Sprintf("%s %s %s %s", m.Name, m.dbType, primary, m.Default)
}

type MetaTable struct {
	Name    string        `json:"Name"`
	Columns []*MetaColumn `json:"columns"`
}

func NewMetaTable(buf []byte) *MetaTable {
	ret := &MetaTable{}
	if err := jsoniter.ConfigFastest.Unmarshal(buf, ret); err != nil {
		panic(err)
	}

	for i := range ret.Columns {
		switch strings.ToLower(ret.Columns[i].Name) {
		case "id":
			ret.Columns[i].Primary = true
		case "create_time", "update_time":
			ret.Columns[i].Default = "default extract(epoch from now())::bigint"
		}

		switch strings.ToLower(ret.Columns[i].Type) {
		case "string":
			ret.Columns[i].dbType = "varchar"
		case "int64":
			ret.Columns[i].dbType = "bigint"
		case "int":
			ret.Columns[i].dbType = "int"
		case "geometry":
			ret.Columns[i].dbType = "geometry(Geometry, 4326)"
		}
	}
	return ret
}

func (m *MetaTable) ColumnsString() string {
	arr := make([]string, 0, len(m.Columns))
	for i := range m.Columns {
		arr = append(arr, m.Columns[i].Name)
	}
	return strings.Join(arr, ",")
}

func (m *MetaTable) ColumnLength() int {
	arr := make([]string, 0, len(m.Columns))
	for i := range m.Columns {
		arr = append(arr, m.Columns[i].Name)
	}
	return len(arr)
}

func (m *MetaTable) CreateTable() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("create table if not exists %s", m.Name))
	buf.WriteString("(\n")
	for i := range m.Columns {
		buf.WriteString(m.Columns[i].ColumnDDL())
		if i < len(m.Columns)-1 {
			buf.WriteByte(',')
		}
		buf.WriteByte('\n')
	}
	buf.WriteString(");\n")
	return buf.String()
}

func (m *MetaTable) PrimaryColumn() *MetaColumn {
	for i := range m.Columns {
		if m.Columns[i].Primary {
			return m.Columns[i]
		}
	}
	return nil
}

func (m *MetaTable) OnConflictDDL() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("on conflict (%s) do update set\n", m.PrimaryColumn().Name))
	for i := range m.Columns {
		if !m.Columns[i].Primary {

			if m.Columns[i].Name == "update_time" {
				buf.WriteString(fmt.Sprintf("%s=GREATEST(%s.%s, excluded.%s)", m.Columns[i].Name, m.Name, m.Columns[i].Name, m.Columns[i].Name))
			} else {
				buf.WriteString(fmt.Sprintf("%s=excluded.%s", m.Columns[i].Name, m.Columns[i].Name))
			}
			if i < len(m.Columns)-1 {
				buf.WriteByte(',')
			} else {
				buf.WriteByte(';')
			}
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}

func (m *MetaTable) Upsert(echos ...*svc.Echo) (string, []interface{}) {
	size := len(echos)
	values := make([]string, 0, size)
	args := make([]interface{}, 0, size*m.ColumnLength())
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
		values = append(values, utils.ValuesPlaceHolder(i*m.ColumnLength(), m.ColumnLength()))
		args = append(args, v.Id, v.BoundingAreaId, v.AccessKey, v.GeofenceId, v.FloorId, v.Floor, v.RoomId,
			strings.Join(v.Macs, ","), createTime, updateTime)
	}

	query := fmt.Sprintf(`insert into %s (%s) values %s %s`, m.Name, m.ColumnsString(),
		strings.Join(values, ","), m.OnConflictDDL())
	return query, args
}

func (m *MetaTable) List(in *svc.ListEchoRequest) (string, []interface{}) {
	firstCond := true
	var buffer bytes.Buffer
	if in.Header {
		buffer.WriteString(fmt.Sprintf("select count(1) from %s", m.Name))
	} else {
		buffer.WriteString(fmt.Sprintf("select * from %s", m.Name))
	}

	if len(in.AccessKeys) > 0 {
		query := fmt.Sprintf("%s access_key in (%s)", utils.CondSql(firstCond), utils.ArraySqlIn(in.AccessKeys...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.BoundingAreaIds) > 0 {
		query := fmt.Sprintf("%s bounding_area_id in (%s)", utils.CondSql(firstCond), utils.ArraySqlIn(in.BoundingAreaIds...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.GeofenceIds) > 0 {
		query := fmt.Sprintf("%s geofence_id in (%s)", utils.CondSql(firstCond), utils.ArraySqlIn(in.GeofenceIds...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.RoomIds) > 0 {
		query := fmt.Sprintf("%s room_id in (%s)", utils.CondSql(firstCond), utils.ArraySqlIn(in.RoomIds...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.FloorsIds) > 0 {
		query := fmt.Sprintf("%s floor_id in (%s)", utils.CondSql(firstCond), utils.ArraySqlIn(in.FloorsIds...))
		buffer.WriteString(query)
		firstCond = false
	}

	if len(in.Floors) > 0 {
		query := fmt.Sprintf("%s floor in (%s)", utils.CondSql(firstCond), utils.ArraySqlIn(in.Floors...))
		buffer.WriteString(query)
		firstCond = false
	}

	if in.EndTime > 0 {
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

func (m *MetaTable) Delete(ids ...string) (string, []interface{}) {
	query := fmt.Sprintf("delete from %s where %s in (%s);", m.Name, m.PrimaryColumn().Name, utils.ArraySqlIn(ids...))
	return query, nil
}

func (m *MetaTable) Get(ids ...string) (string, []interface{}) {
	query := fmt.Sprintf("select * from %s where %s in (%s);", m.Name, m.PrimaryColumn().Name, utils.ArraySqlIn(ids...))
	return query, nil
}
