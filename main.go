package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vx416/modelgen/pkg/modelgen"

	"github.com/vx416/modelgen/pkg/setting"
)

var s = &setting.Settings{}

func newSettings() {
	flag.StringVar(&s.InputPath, "i", "", "file path of schema file")
	flag.StringVar(&s.TableNames, "tables", "", "table name splited by comma")
	flag.StringVar(&s.OutputPath, "o", "", "destination of output content")
	flag.StringVar(&s.PbOutputPath, "pbo", "", "destination of output protocol buffer content")
	flag.StringVar(&s.Tag, "tag", "db", "model tag")
	flag.StringVar(&s.PackageName, "package", "model", "package of model")
	flag.StringVar(&s.DBKind, "db", "mysql", "kind of database")

	flag.BoolVar(&s.Print, "print", true, "print out of generated content")
	flag.BoolVar(&s.Pb, "pb", true, "print out of protocol buffer content")
	flag.BoolVar(&s.JsonTag, "json", true, "add json tag")
	flag.BoolVar(&s.AppendOnly, "append", false, "append only")
	flag.Parse()
}

func main() {
	newSettings()

	models, err := modelgen.FilterModelsFromPath(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := modelgen.GetOutput(s.AppendOnly, s.PackageName, models)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if s.Print {
		fmt.Println(res)
		fmt.Println(strings.Repeat("=", 30))
	}
	if s.OutputPath != "" {
		outputFile, err := os.Create(s.OutputPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputFile.WriteString(res)
	}

	if s.ModelSettings.Pb {
		res, err := modelgen.GetPbOutput(s.AppendOnly, s.PackageName, models)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(res)
		fmt.Println(strings.Repeat("=", 30))

		if s.PbOutputPath != "" {
			outputFile, err := os.Create(s.PbOutputPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			outputFile.WriteString(res)
		}
	}
}
