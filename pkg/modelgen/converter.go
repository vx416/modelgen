package modelgen

import (
	"fmt"
	"strings"

	"github.com/vx416/modelgen/pkg/dbhelper"
	"github.com/vx416/modelgen/pkg/setting"
	"github.com/xwb1989/sqlparser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	caser = cases.Title(language.English, cases.NoLower)
)

func Converter(columns []*sqlparser.ColumnDefinition, setting *setting.Settings) ([]*StructField, error) {
	sfs := make([]*StructField, 0, len(columns))
	for _, col := range columns {
		tag := setting.GetTag(col.Name.String())
		sf, err := convertToStructField(setting.GetHelper(), col, tag)
		if err != nil {
			return nil, err
		}
		sfs = append(sfs, &sf)
	}
	return sfs, nil
}

type StructField struct {
	Name string
	Type string
	Tag  string
}

func (sf StructField) String() string {
	var structFields strings.Builder
	structFields.WriteString(sf.Name)
	structFields.WriteString(" ")
	structFields.WriteString(sf.Type)
	structFields.WriteString(" ")
	structFields.WriteString(sf.Tag)
	return structFields.String()
}

func convertToStructField(helper dbhelper.Helper, column *sqlparser.ColumnDefinition, tag string) (StructField, error) {
	colType := column.Type.Type
	stField := StructField{
		Name: camelCaseString(column.Name.String()),
		Tag:  tag,
	}
	if helper.IsString(colType) || helper.IsText(colType) {
		stField.Type = "string"
		if !column.Type.NotNull {
			stField.Type = "null.String"
		}
	}
	if helper.IsInteger(colType) {
		stField.Type = "int64"
		if !column.Type.NotNull {
			stField.Type = "null.Int"
		}
	}
	if helper.IsSmallInterger(colType) {
		stField.Type = "int8"
		if !column.Type.NotNull {
			stField.Type = "null.Int"
		}
	}
	if helper.IsFloat(colType) {
		stField.Type = "decimal.Decimal"
		if !column.Type.NotNull {
			stField.Type = "decimal.NullDecimal"
		}
	}
	if helper.IsTimestamp(colType) {
		stField.Type = "time.Time"
		if !column.Type.NotNull {
			stField.Type = "null.Time"
		}
	}
	if stField.Type == "" {
		switch colType {
		case "boolean":
			stField.Type = "bool"
			if !column.Type.NotNull {
				stField.Type = "null.Bool"
			}
		default:
			stField.Type = "null.String"
		}
	}

	return stField, nil
}

func camelCaseString(s string) string {
	if s == "" {
		return s
	}
	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		if s == "id" {
			return "ID"
		}
		return caser.String(s)
	}

	var cc string
	for _, part := range splitted {
		if part == "id" {
			cc += "ID"
			continue
		}
		cc += caser.String(strings.ToLower(part))
	}
	return cc
}

func dbTag(colName string) string {
	return fmt.Sprintf(`db:"%s"`, colName)
}

func gormTag(colName string) string {
	return fmt.Sprintf(`gorm:"column:%s"`, colName)
}
