package meta

import (
	"bytes"
	"fmt"
	"github.com/linger1216/go-front/utils"
)

type TableInfo struct {
	Name     string
	Comment  string
	IndexDDL []string
	Columns  []*ColumnInfo
}

func (c *TableInfo) ToGolangDefinition(annotation bool) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("type %s struct {\n", utils.Case2Camel(c.Name)))
	for _, v := range c.Columns {
		buf.WriteString(v.toGolangDefinition(annotation))
		buf.WriteByte('\n')
	}
	buf.WriteString("}\n")
	return buf.String()
}
