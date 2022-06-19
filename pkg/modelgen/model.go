package modelgen

import (
	"github.com/pkg/errors"

	"github.com/vx416/modelgen/pkg/setting"
	"github.com/vx416/modelgen/pkg/util"
)

type Model struct {
	Name      string
	TableName string
	Fields    []string
	PbFields  []string
}

func NewModels(setting *setting.Settings, ddlStr string) ([]*Model, error) {
	ddls, err := Parse(ddlStr)
	if err != nil {
		return nil, errors.Wrap(err, "parser ddl failed")
	}

	models := make([]*Model, 0, len(ddls))

	for _, ddl := range ddls {
		model := Model{
			Name:      util.Singular(util.CamelCaseString(ddl.TableName)),
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
		if setting.Pb {
			pbs, err := PbConverter(ddl.Columns, setting)
			if err != nil {
				return nil, err
			}
			model.PbFields = make([]string, 0, len(pbs))
			for _, pb := range pbs {
				model.PbFields = append(model.PbFields, pb.String())
			}
		}
	}

	return models, nil
}
