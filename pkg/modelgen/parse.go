package modelgen

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"github.com/vx416/modelgen/pkg/setting"
	"github.com/xwb1989/sqlparser"
)

var (
	ErrDDLInvalid = errors.New("statement is not DDL string")
)

func FilterModelsFromPath(s *setting.Settings) ([]*Model, error) {
	tableNames := strings.Split(s.TableNames, ";")
	tableNamesMap := make(map[string]bool)
	for _, t := range tableNames {
		tableNamesMap[strings.ToLower(t)] = true
	}
	filterModels := make([]*Model, 0, 10)

	inputInfo, err := os.Stat(s.InputPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("input(%s) not exists\n", s.InputPath)
			os.Exit(1)
		}
		fmt.Printf("input(%s) invalid\n", s.InputPath)
		os.Exit(1)
	}

	if inputInfo.IsDir() {
		err := filepath.Walk(s.InputPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				models, err := NewModels(s, string(data))
				if err != nil {
					return err
				}
				for _, m := range models {
					if tableNamesMap[strings.ToLower(m.TableName)] {
						filterModels = append(filterModels, m)
					}
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return filterModels, nil
}

func Parse(ddl string) ([]TableDDL, error) {
	ss := strings.Split(ddl, ";")
	ddls := make([]TableDDL, 0, 10)

	for _, s := range ss {
		if strings.TrimSpace(s) == "" {
			continue
		}
		if strings.Contains(s, "CREATE TABLE") || strings.Contains(s, "create table") {
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
