package modelgen

import (
	"bytes"
	"html/template"
	"strings"
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

{{- range .Methods}}
{{.}}
{{- end}}
{{end}}
`

const pbModelTemptae = `
{{if .Append | not -}}
syntax = "proto3";

package {{.PackageName}}

import "google/protobuf/timestamp.proto";
import "google/protobuf/empty.proto";
{{ end -}}

{{ range .Models }}
message {{.Name}} {
	{{- range $index, $field := .PbFields}}
	{{$field}} = {{$index | AddOne}};
	{{- end}}
}
{{end}}

`

func GetOutput(appendOnly bool, packageName string, models []*Model) (string, error) {
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

func GetPbOutput(appendOnly bool, packageName string, models []*Model) (string, error) {
	t, err := template.New("").Funcs(map[string]any{
		"AddOne": addOne,
	}).Parse(pbModelTemptae)
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

func addOne(i int) int {
	return i + 1
}
