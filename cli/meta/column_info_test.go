package meta

import "testing"

func TestColumnInfo_toGolangMemberDef(t *testing.T) {
	col := &ColumnInfo{
		Name:    "geofence_id",
		Type:    "character varying",
		Comment: "地理围栏id",
		Default: "",
		Primary: false,
	}

	col.toGolangDefinition(true)
}
