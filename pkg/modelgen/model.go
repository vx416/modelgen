package modelgen

import (
	"github.com/pkg/errors"

	"github.com/vx416/modelgen/pkg/setting"
)

type Model struct {
	Name      string
	TableName string
	Fields    []string
}

func NewModels(setting *setting.Settings, ddlStr string) ([]*Model, error) {
	ddls, err := Parse(ddlStr)
	if err != nil {
		return nil, errors.Wrap(err, "parser ddl failed")
	}

	models := make([]*Model, 0, len(ddls))

	for _, ddl := range ddls {
		model := Model{
			Name:      camelCaseString(ddl.TableName),
			TableName: ddl.TableName,
		}
		sfs, err := Converter(ddl.Columns, setting)
		if err != nil {
			return nil, err
		}
		model.Fields = make([]string, 0, len(sfs))
		for _, sf := range sfs {
			model.Fields = append(model.Fields, sf.String())
		}

		models = append(models, &model)
	}

	return models, nil
}
