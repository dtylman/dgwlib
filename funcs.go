package dgwlib

import (
	"fmt"
	"text/template"
)

var tmplFuncMap = template.FuncMap{
	"createInsertSQL":            CreateInsertSQL,
	"createInsertParams":         CreateInsertParams,
	"createInsertScan":           CreateInsertScan,
	"createSelectByPkSQL":        CreateSelectByPkSQL,
	"createSelectByPkFuncParams": CreateSelectByPkFuncParams,
	"createSelectByPkSQLParams":  CreateSelectByPkSQLParams,
	"createSelectByPkScan":       CreateSelectByPkScan,
}

// CreateSelectByPkSQL ...
func CreateSelectByPkSQL(st *Struct) string {
	var sql string
	var colNames []string
	var pkNames []string
	for _, c := range st.Table.Columns {
		if c.IsPrimaryKey {
			pkNames = append(pkNames, c.Name)
		}
		colNames = append(colNames, c.Name)
	}
	sql = "SELECT " + flatten(colNames, ", ") + " FROM " + st.Table.Name + " WHERE "
	for i, c := range pkNames {
		placeHolder := i + 1
		if i == 0 {
			sql = sql + c + fmt.Sprintf(" = $%d", placeHolder)
		} else {
			sql = sql + " AND " + c + fmt.Sprintf(" = $%d", placeHolder)
		}
	}
	return sql
}

// CreateSelectByPkScan ...
func CreateSelectByPkScan(st *Struct) string {
	var s []string
	for _, f := range st.Fields {
		s = append(s, fmt.Sprintf("&r.%s", f.Name))
	}
	return flatten(s, ", ")
}

// CreateSelectByPkSQLParams ...
func CreateSelectByPkSQLParams(st *Struct) string {
	var fs []string
	for i, f := range st.Fields {
		if f.Column.IsPrimaryKey {
			fs = append(fs, fmt.Sprintf("pk%d", i))
		}
	}
	return flatten(fs, ", ")
}

// CreateSelectByPkFuncParams ...
func CreateSelectByPkFuncParams(st *Struct) string {
	var fs []string
	for i, f := range st.Fields {
		if f.Column.IsPrimaryKey {
			fs = append(fs, fmt.Sprintf("pk%d ", i)+f.Type)
		}
	}
	return flatten(fs, ", ")
}

// CreateInsertScan ...
func CreateInsertScan(st *Struct) string {
	var fs []string
	for _, f := range st.Fields {
		if f.Column.IsPrimaryKey && st.Table.AutoGenPk {
			fs = append(fs, "&r."+f.Name)
		}
	}
	return flatten(fs, ", ")
}

// CreateInsertParams ...
func CreateInsertParams(st *Struct) string {
	var fs []string
	for _, f := range st.Fields {
		if f.Column.IsPrimaryKey && st.Table.AutoGenPk {
			continue
		} else {
			fs = append(fs, "&r."+f.Name)
		}
	}
	return flatten(fs, ", ")
}

func flatten(elems []string, sep string) string {
	var str string
	for i, e := range elems {
		if i == 0 {
			str = str + e
		} else {
			str = str + sep + e
		}
	}
	return str
}

func placeholders(l []string) string {
	var ph string
	var j int
	for i := range l {
		j = i + 1
		if i == 0 {
			ph = ph + fmt.Sprintf("$%d", j)
		} else {
			ph = ph + fmt.Sprintf(", $%d", j)
		}
	}
	return ph
}

// CreateInsertSQL ...
func CreateInsertSQL(st *Struct) string {
	var sql string
	sql = "INSERT INTO " + st.Table.Name + " ("

	if len(st.Table.Columns) == 1 && st.Table.Columns[0].IsPrimaryKey && st.Table.AutoGenPk {
		sql = sql + st.Table.Columns[0].Name + ") VALUES (DEFAULT)"
	} else {
		var colNames []string
		for _, c := range st.Table.Columns {
			if c.IsPrimaryKey && st.Table.AutoGenPk {
				continue
			} else {
				colNames = append(colNames, c.Name)
			}
		}
		sql = sql + flatten(colNames, ", ") + ") VALUES ("

		var fieldNames []string
		for _, f := range st.Fields {
			if f.Column.IsPrimaryKey && st.Table.AutoGenPk {
				continue
			} else {
				fieldNames = append(fieldNames, f.Name)
			}
		}
		sql = sql + placeholders(fieldNames) + ")"
	}

	if st.Table.AutoGenPk {
		sql = sql + " RETURNING "
		for i, c := range st.Table.PrimaryKeys {
			if i == 0 {
				sql = sql + c.Name
			} else {
				sql = sql + ", " + c.Name
			}
		}
	}
	return sql
}
