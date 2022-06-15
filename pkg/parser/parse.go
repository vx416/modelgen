package parser

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xwb1989/sqlparser"
)

var (
	ErrDDLInvalid = errors.New("statement is not DDL string")
)

func Parse(ddl string) ([]TableDDL, error) {
	ss := strings.Split(ddl, ";")
	ddls := make([]TableDDL, 0, 10)

	for _, s := range ss {
		if strings.TrimSpace(s) == "" {
			continue
		}
		ddl, err := ParseDDL(s)
		if err != nil && errors.Is(err, ErrDDLInvalid) {
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		ddls = append(ddls, ddl)
	}
	return ddls, nil
}

func ParseDDL(ddlStr string) (TableDDL, error) {
	stmt, err := sqlparser.Parse(ddlStr)
	if err != nil {
		return TableDDL{}, errors.Wrapf(err, "parse sql failed, %s", ddlStr)
	}
	ddlStmt, ok := stmt.(*sqlparser.DDL)
	if !ok {
		return TableDDL{}, ErrDDLInvalid
	}
	tableName := ddlStmt.NewName.Name.String()
	if tableName == "" {
		return TableDDL{}, errors.New("table name cannot be empty")
	}
	columns := ddlStmt.TableSpec.Columns
	if len(columns) == 0 {
		return TableDDL{}, errors.New("column cannot be empty")
	}

	return TableDDL{
		TableName: tableName,
		Columns:   columns,
	}, nil
}

type TableDDL struct {
	TableName string
	Columns   []*sqlparser.ColumnDefinition
}
