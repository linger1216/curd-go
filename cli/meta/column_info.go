package meta

import (
	"bytes"
	"fmt"
	"github.com/linger1216/go-front/utils"
)

type ColumnInfo struct {
	Name    string
	Type    string
	Comment string
	Default string
	Primary bool
}

func (c *ColumnInfo) toGolangDefinition(annotation bool) string {
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
