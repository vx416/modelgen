package output

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/vx416/modelgen/pkg/modelgen"
)

const modelTemplate = `
{{if .Append | not -}}
package {{.PackageName}}

import (
	"gopkg.in/guregu/null.v4"	
	"github.com/shopspring/decimal"
)
{{ end -}}

{{ range .Models }}
type {{.Name}} struct {
	{{- range .Fields}}
	{{.}}
	{{- end}}
}
{{end}}
`

func Output(appendOnly bool, packageName string, models []*modelgen.Model) (string, error) {
	t, err := template.New("").Parse(modelTemplate)
	if err != nil {
		return "", err
	}
	var buf = bytes.NewBufferString("")
	err = t.Execute(buf, map[string]interface{}{
		"Append":      appendOnly,
		"PackageName": packageName,
		"Models":      models,
	})
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(buf.String(), "&#34;", `"`), nil
}
