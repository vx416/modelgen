package setting

import (
	"fmt"

	"github.com/vx416/modelgen/pkg/dbhelper"
)

type Settings struct {
	ModelSettings
	OutputSettings
	InputPath string
}

type OutputSettings struct {
	Print       bool
	Destination string
	PackageName string
	AppendOnly  bool
}

type ModelSettings struct {
	Tag    string `db:"id"`
	DBKind string
}

func (set ModelSettings) GetTag(col string) string {
	switch set.Tag {
	case "db", "DB":
		return dbTag(col)
	case "gorm", "GORM":
		return gormTag(col)
	default:
		return dbTag(col)
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

func dbTag(colName string) string {
	return "`db:" + fmt.Sprintf(`"%s"`, colName) + "`"
}

func gormTag(colName string) string {
	return "`gorm:column" + `"` + colName + `"` + "`"
}
