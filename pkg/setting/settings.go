package setting

import (
	"fmt"

	"github.com/vx416/modelgen/pkg/dbhelper"
	"github.com/vx416/modelgen/pkg/util"
)

type Settings struct {
	ModelSettings
	OutputSettings
	InputPath string
}

type OutputSettings struct {
	Print        bool
	OutputPath   string
	PbOutputPath string
	PackageName  string
	AppendOnly   bool
}

type ModelSettings struct {
	Pb         bool
	Tag        string
	JsonTag    bool
	TableNames string
	DBKind     string
}

func (set ModelSettings) GetTag(col string) string {
	switch set.Tag {
	case "db", "DB":
		return dbTag(col, set.JsonTag)
	case "gorm", "GORM":
		return gormTag(col, set.JsonTag)
	default:
		return dbTag(col, set.JsonTag)
	}
}

func (set ModelSettings) GetHelper() dbhelper.Helper {
	switch set.DBKind {
	case "mysql", "MySQL", "MYSQL":
		return dbhelper.MySQL{}
	default:
		return dbhelper.MySQL{}
	}
}

func dbTag(colName string, jsonTagOn bool) string {
	if jsonTagOn {
		return fmt.Sprintf("`%s db:%s%s%s`", jsonTag(colName), `"`, colName, `"`)
	}
	return fmt.Sprintf("`db:%s%s%s`", `"`, colName, `"`)
}

func gormTag(colName string, jsonTagOn bool) string {
	if jsonTagOn {
		return fmt.Sprintf("`%s gorm:%scolumn:%s%s`", jsonTag(colName), `"`, colName, `"`)
	}
	return fmt.Sprintf("`gorm:%scolumn:%s%s`", `"`, colName, `"`)
}

func jsonTag(colName string) string {
	return fmt.Sprintf(`json:"%s"`, util.LowercaseCamelCaseString(colName))
}
