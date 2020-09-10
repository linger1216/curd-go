package meta

const (
	ShowTablesQuery = `select relname as table_name,cast(obj_description(relfilenode,'pg_class') as varchar) as comment from pg_class c 
where  relkind = 'r' and relname not like 'pg_%' and relname not like 'sql_%' and relchecks=0 order by relname;`
	ShowTableColumnIndexQuery = `select tablename as table_name , indexname as index_name, indexdef as ddl from pg_indexes where tablename='%s';`
	ShowTableColumnQuery      = `
SELECT DISTINCT
    a.attnum as num,
    a.attname as name,
    format_type(a.atttypid, a.atttypmod) as type,
    a.attnotnull as notnull, 
    com.description as comment,
    coalesce(i.indisprimary,false) as primary_key,
    def.adsrc as default
FROM pg_attribute a 
JOIN pg_class pgc ON pgc.oid = a.attrelid
LEFT JOIN pg_index i ON 
    (pgc.oid = i.indrelid AND i.indkey[0] = a.attnum)
LEFT JOIN pg_description com on 
    (pgc.oid = com.objoid AND a.attnum = com.objsubid)
LEFT JOIN pg_attrdef def ON 
    (a.attrelid = def.adrelid AND a.attnum = def.adnum)
WHERE a.attnum > 0 AND pgc.oid = a.attrelid
AND pg_table_is_visible(pgc.oid)
AND NOT a.attisdropped
AND pgc.relname = '%s'
ORDER BY a.attnum;
`
)

type ColumnInfo struct {
	Name    string
	Type    string
	Comment string
	Default string
	Primary bool
}

type TableInfo struct {
	Name     string
	Comment  string
	IndexDDL []string
	Columns  []*ColumnInfo
}
