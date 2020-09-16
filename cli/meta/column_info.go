package meta

import (
	"bytes"
	"fmt"
	"github.com/linger1216/go-front/utils"
	"strings"
)

type ColumnInfo struct {
	Name    string
	Type    string
	Comment string
	Default string
	Primary bool
}

func (c *ColumnInfo) GolangModelDefinition(annotation bool) string {
	var buf bytes.Buffer
	camelName := utils.Case2Camel(c.Name)
	t := pgType2GoType(c.Type)
	buf.WriteString(fmt.Sprintf("    %s %s ", camelName, t))
	if annotation {
		buf.WriteString("`")
		buf.WriteString(fmt.Sprintf(`form:"%s" json:"%s,omitempty" yaml:"%s"`, c.Name,
			utils.LowerFirst(camelName), utils.LowerFirst(camelName)))
		buf.WriteString("`")
	}
	if len(c.Comment) > 0 {
		buf.WriteString(fmt.Sprintf(" // %s", c.Comment))
	}
	return buf.String()
}

func (c *ColumnInfo) DBSelectColumn() string {
	switch c.Type {
	case `character varying`:
		return c.Name
	case `bigint`:
		return c.Name
	case `integer`:
		return c.Name
	case `geometry(Geometry,4326)`:
		return fmt.Sprintf("st_asgeojson(%s) as %s", c.Name, c.Name)
	case `character varying[]`:
		return fmt.Sprintf("array_to_string(%s, ',') as %s", c.Name, c.Name)
	case `integer[]`:
		return fmt.Sprintf("array_to_string(%s, ',') as %s", c.Name, c.Name)
	}
	return ""
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

func (c *ColumnInfo) DBCreateTableDDL(annotation bool) string {
	var primary string
	if c.Primary {
		primary = "primary key"
	}
	return fmt.Sprintf("%s %s %s %s", c.Name, c.Type, primary, c.Default)
}

// 赋值的具体过程
func (c *ColumnInfo) DBUpsertFormat() string {
	if c.Primary {
		return fmt.Sprintf(`
		if len(v.%s) == 0 {
			v.%s = utils.Generate()
		}`, c.Name, c.Name)
	}

	switch c.Name {
	case "create_time":
		return fmt.Sprintf(`
		var createTime int64
    if v.CreateTime == 0 {
			createTime = time.Now().Unix()
		} else {
			createTime = v.CreateTime
		}`)
	case "update_time":
		return fmt.Sprintf(`
		var updateTime int64
    if v.UpdateTime == 0 {
			updateTime = time.Now().Unix()
		} else {
			updateTime = v.UpdateTime
		}`)
	}

	if c.Type == "geometry(Geometry,4326)" {
		return fmt.Sprintf(`
		geometry, err := jsoniter.ConfigFastest.Marshal(v.Geometry)
		if err != nil {
			panic(err)
		}`)
	}

	return ""
}

// 最后赋值的变量
func (c *ColumnInfo) DBUpsertVariable() string {
	CaseName := utils.Case2Camel(c.Name)
	switch c.Name {
	case "create_time":
		return "create_time"
	case "update_time":
		return "update_time"
	}
	switch c.Type {
	case "geometry(Geometry,4326)":
		return "geometry"
	case "character varying[]", "integer[]":
		return fmt.Sprintf("pq.Array(v.%s)", CaseName)
	}
	return fmt.Sprintf("v.%s", CaseName)
}

func (c *ColumnInfo) DBListFormat() string {
	switch strings.ToLower(c.Name) {
	case "update_time":
		return `	
    if in.StartTime > 0 && in.EndTime > 0 {
		query := fmt.Sprintf("%s update_time between '%d' and '%d' ", utils.CondSql(firstCond), in.StartTime, in.EndTime)
		buffer.WriteString(query)
		firstCond = false
		}`
	}

	arrContain := func(name, funcName string) string {
		CaseName := utils.Case2Camel(c.Name)
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("if len(in.%s) > 0 {\n", CaseName))
		buf.WriteString(`query := fmt.Sprintf("%s ` + c.Name + ` @> %s", utils.CondSql(firstCond), utils.` + funcName + `(in.` + CaseName + "...))")
		buf.WriteString("buffer.WriteString(query)\n")
		buf.WriteString("firstCond = false\n")
		buf.WriteString("}\n")
		return buf.String()
	}

	oneIn := func(name, funcName string) string {
		CaseName := utils.Case2Camel(c.Name)
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("if len(in.%ss) > 0 {\n", CaseName))
		buf.WriteString(`query := fmt.Sprintf("%s ` + c.Name + ` in (%s)", utils.CondSql(firstCond), utils.` + funcName + `(in.` + CaseName + "...))")
		buf.WriteString("buffer.WriteString(query)\n")
		buf.WriteString("firstCond = false\n")
		buf.WriteString("}\n")
		return buf.String()
	}

	switch c.Type {
	case "geometry(Geometry,4326)":
		return `	
		if in.Point != nil && in.Radius > 0 {
		query := fmt.Sprintf("%s %s", utils.CondSql(firstCond),
			utils.SqlWithIn(in.Point.Coordinates[0], in.Point.Coordinates[1], int(in.Radius)))
		buffer.WriteString(query)
		firstCond = false
	  }`
	case "character varying[]":
		return arrContain(c.Name, "SqlStringArray")
	case "integer[]":
		return arrContain(c.Name, "SqlIntegerArray")
	case "integer":
		return oneIn(c.Name, "SqlIntegerIn")
	case "character varying":
		return oneIn(c.Name, "SqlStringIn")
	}

	return ""
}
