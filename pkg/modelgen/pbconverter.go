package modelgen

import (
	"strings"

	"github.com/vx416/modelgen/pkg/dbhelper"
	"github.com/vx416/modelgen/pkg/setting"
	"github.com/vx416/modelgen/pkg/util"
	"github.com/xwb1989/sqlparser"
)

func PbConverter(columns []*sqlparser.ColumnDefinition, setting *setting.Settings) ([]*PbStructField, error) {
	sfs := make([]*PbStructField, 0, len(columns))
	for _, col := range columns {
		sf, err := convertToPbStructField(setting.GetHelper(), col)
		if err != nil {
			return nil, err
		}
		sfs = append(sfs, &sf)
	}
	return sfs, nil
}

type PbStructField struct {
	Name string
	Type string
}

func (sf PbStructField) String() string {
	var structFields strings.Builder
	structFields.WriteString(sf.Type)
	structFields.WriteString(" ")
	structFields.WriteString(sf.Name)
	return structFields.String()
}

func convertToPbStructField(helper dbhelper.Helper, column *sqlparser.ColumnDefinition) (PbStructField, error) {
	colType := column.Type.Type
	stField := PbStructField{
		Name: util.LowercaseCamelCaseString(column.Name.String()),
	}
	if helper.IsString(colType) || helper.IsText(colType) {
		stField.Type = "string"
	}
	if helper.IsInteger(colType) {
		stField.Type = "int64"
	}
	if helper.IsSmallInterger(colType) {
		stField.Type = "int32"
	}
	if helper.IsFloat(colType) {
		stField.Type = "float"
	}
	if helper.IsTimestamp(colType) {
		stField.Type = "google.protobuf.Timestamp"
	}
	if stField.Type == "" {
		switch colType {
		case "boolean":
			stField.Type = "bool"
		default:
			stField.Type = "string"
		}
	}

	return stField, nil
}
