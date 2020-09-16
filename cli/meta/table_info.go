package meta

import (
	"bytes"
	"fmt"
	"github.com/linger1216/go-front/utils"
	"strings"
)

type TableInfo struct {
	Name     string
	Comment  string
	IndexDDL []string
	Columns  []*ColumnInfo
}

func (c *TableInfo) GolangModelDefinition(annotation bool) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("type %s struct {\n", utils.Case2Camel(c.Name)))
	for _, v := range c.Columns {
		buf.WriteString(v.GolangModelDefinition(annotation))
		buf.WriteByte('\n')
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (c *TableInfo) DBSelectColumn() string {
	var buf bytes.Buffer
	for i := range c.Columns {
		buf.WriteString(c.Columns[i].DBSelectColumn())
		if i < len(c.Columns)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func (c *TableInfo) DBColumns() string {
	var buf bytes.Buffer
	for i := range c.Columns {
		buf.WriteString(c.Columns[i].Name)
		if i < len(c.Columns)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

/*
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
*/
func (c *TableInfo) DBCreateTableDDL(annotation bool) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("create table if not exists %s\n", c.Name))
	buf.WriteString("(\n")
	for i, v := range c.Columns {
		buf.WriteString(fmt.Sprintf("%s", v.DBCreateTableDDL(annotation)))
		if i < len(c.Columns)-1 {
			buf.WriteString(",")
		}
		buf.WriteByte('\n')
	}
	buf.WriteString(");\n")
	return buf.String()
}

func (c *TableInfo) DBIndexTableDDL() []string {
	arr := make([]string, 0)
	for i := range c.IndexDDL {
		ddl := strings.ToLower(c.IndexDDL[i])
		pos := strings.Index(ddl, "index")
		arr = append(arr, ddl[:pos+5]+"if not exists"+ddl[pos+5:])
	}
	return arr
}

/*
on conflict (id)
do update set
age = excluded.age,
name = excluded.name,
geometry = excluded.geometry,
books = excluded.books,
tags = excluded.tags,
update_time = GREATEST(%s.update_time, excluded.update_time);
*/
func (c *TableInfo) DBOnConflictDDL() string {
	primaryColumn := c.DBPrimaryColumn()
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("on conflict (%s)\n", primaryColumn.Name))
	buf.WriteString("do update set\n")

	for i, v := range c.Columns {
		if v.Primary || v.Name == "create_time" {
			continue
		}
		if v.Name == "update_time" {
			buf.WriteString(fmt.Sprintf("update_time = GREATEST(%s.update_time, excluded.update_time)", c.Name))
		} else {
			buf.WriteString(fmt.Sprintf("%s = excluded.%s", v.Name, v.Name))
		}
		if i < len(c.Columns)-1 {
			buf.WriteString(",")
		} else {
			buf.WriteString(";")
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (c *TableInfo) DBPrimaryColumn() *ColumnInfo {
	for i := range c.Columns {
		if c.Columns[i].Primary {
			return c.Columns[i]
		}
	}
	return nil
}

func (c *TableInfo) DBUpsert() string {
	var buf bytes.Buffer

	buf.WriteString(`
	cols := strings.Split(m.ColumnsString(), ",")
	size := len(echos)
	values := make([]string, 0, size)
	args := make([]interface{}, 0, size*len(cols))
	for i, v := range echos {`)
	buf.WriteByte('\n')

	hasGeometry := false
	geometryIndex := 0
	for i := range c.Columns {
		if c.Columns[i].Type == "geometry(Geometry,4326)" {
			hasGeometry = true
			geometryIndex = i + 1
		}
		buf.WriteString(c.Columns[i].DBUpsertFormat())
		buf.WriteByte('\n')
	}

	if hasGeometry {
		buf.WriteString(fmt.Sprintf(`values = append(values, utils.ValuePlaceHolderAndGeometry(i*len(cols), len(cols), %d))`, geometryIndex))
	} else {
		buf.WriteString(fmt.Sprintf(`values = append(values, utils.ValuePlaceHolder(i*len(cols), len(cols)))`))
	}
	buf.WriteByte('\n')

	args := make([]string, 0)
	for i := range c.Columns {
		args = append(args, c.Columns[i].DBUpsertVariable())
	}

	buf.WriteString(fmt.Sprintf("args = append(args, %s)", strings.Join(args, ",")))
	buf.WriteByte('\n')
	buf.WriteString("}\n")

	buf.WriteString(`	query := fmt.Sprintf("insert into %s (%s) values %s %s", m.Table(), m.ColumnsString(),
		strings.Join(values, ","), m.OnConflictDDL())
	return query, args`)

	return buf.String()
}

func (c *TableInfo) DBList() string {
	var buf bytes.Buffer

	buf.WriteString(`	
	firstCond := true
	var buffer bytes.Buffer
	if in.Header {
		buffer.WriteString(fmt.Sprintf("select count(1) from %s", m.Table()))
	} else {
		buffer.WriteString(fmt.Sprintf("select %s from %s", m.Select(), m.Table()))
	}`)
	buf.WriteString("\n")

	for i := range c.Columns {
		buf.WriteString(c.Columns[i].DBListFormat())
		buf.WriteString("\n")
	}

	buf.WriteString(`	
  if !in.Header {
		query := fmt.Sprintf(" offset %d limit %d", in.CurrentPage*in.PageSize, in.PageSize)
		buffer.WriteString(query)
	}

	buffer.WriteString(";")
	return buffer.String(), nil`)

	return buf.String()
}
