package meta

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/linger1216/go-front/utils"
)

type PostgresMeta struct {
	db *sqlx.DB
}

func NewPostgresMeta(db *sqlx.DB) *PostgresMeta {
	return &PostgresMeta{db}
}

func (p *PostgresMeta) GetInfo() ([]*TableInfo, error) {
	rows, err := p.db.Queryx(ShowTablesQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	ret := make([]*TableInfo, 0)
	for rows.Next() {
		line := make(map[string]interface{})
		err = rows.MapScan(line)
		if err != nil {
			return nil, err
		}
		if tc, err := transTableInfo("", line); err == nil && tc != nil {
			ret = append(ret, tc)
		}
	}

	// 获取字段
	for i := range ret {
		columnInfo, err := p.GetColumnInfo(ret[i].Name)
		if err != nil {
			return nil, err
		}
		ret[i].Columns = columnInfo
	}

	// 获取索引
	for i := range ret {
		indexDDLs, err := p.GetIndexDDL(ret[i].Name)
		if err != nil {
			return nil, err
		}
		ret[i].IndexDDL = indexDDLs
	}
	return ret, nil
}

func (p *PostgresMeta) GetIndexDDL(name string) ([]string, error) {
	rows, err := p.db.Queryx(fmt.Sprintf(ShowTableColumnIndexQuery, name))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	ret := make([]string, 0)
	for rows.Next() {
		line := make(map[string]interface{})
		err = rows.MapScan(line)
		if err != nil {
			return nil, err
		}
		if tc, err := transIndexDDL("", line); err == nil && len(tc) > 0 {
			ret = append(ret, tc)
		}
	}
	return ret, nil
}

func (p *PostgresMeta) GetColumnInfo(name string) ([]*ColumnInfo, error) {
	rows, err := p.db.Queryx(fmt.Sprintf(ShowTableColumnQuery, name))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	ret := make([]*ColumnInfo, 0)
	for rows.Next() {
		line := make(map[string]interface{})
		err = rows.MapScan(line)
		if err != nil {
			return nil, err
		}
		if tc, err := transColumnInfo("", line); err == nil && tc != nil {
			ret = append(ret, tc)
		}
	}
	return ret, nil
}

func transTableInfo(prefix string, m map[string]interface{}) (*TableInfo, error) {
	ret := &TableInfo{}
	if v, ok := m[prefix+"table_name"]; ok {
		ret.Name = utils.ToString(v)
	}
	if v, ok := m[prefix+"comment"]; ok {
		ret.Comment = utils.ToString(v)
	}
	return ret, nil
}

func transColumnInfo(prefix string, m map[string]interface{}) (*ColumnInfo, error) {
	ret := &ColumnInfo{}
	if v, ok := m[prefix+"name"]; ok {
		ret.Name = utils.ToString(v)
	}
	if v, ok := m[prefix+"type"]; ok {
		ret.Type = utils.ToString(v)
	}
	if v, ok := m[prefix+"comment"]; ok {
		ret.Comment = utils.ToString(v)
	}
	if v, ok := m[prefix+"primary_key"]; ok {
		ret.Primary = v.(bool)
	}
	if v, ok := m[prefix+"default"]; ok {
		ret.Default = utils.ToString(v)
	}
	return ret, nil
}

func transIndexDDL(prefix string, m map[string]interface{}) (string, error) {
	if v, ok := m[prefix+"ddl"]; ok {
		return utils.ToString(v), nil
	}
	return "", nil
}
