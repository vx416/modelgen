package modelgen

import (
	"strings"

	"github.com/vx416/modelgen/pkg/dbhelper"
	"github.com/vx416/modelgen/pkg/setting"
	"github.com/vx416/modelgen/pkg/util"
	"github.com/xwb1989/sqlparser"
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
	Name    string
	Type    string
	Tag     string
	Comment string
}

func (sf StructField) String() string {
	var structFields strings.Builder
	structFields.WriteString(sf.Name)
	structFields.WriteString(" ")
	structFields.WriteString(sf.Type)
	structFields.WriteString(" ")
	structFields.WriteString(sf.Tag)
	if sf.Comment != "" {
		structFields.WriteString("// ")
		structFields.WriteString(sf.Comment)
	}

	return structFields.String()
}

func convertToStructField(helper dbhelper.Helper, column *sqlparser.ColumnDefinition, tag string) (StructField, error) {
	colType := column.Type.Type
	stField := StructField{
		Name: util.CamelCaseString(column.Name.String()),
		Tag:  tag,
	}
	if column.Type.Comment != nil {
		stField.Comment = string(column.Type.Comment.Val)
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
