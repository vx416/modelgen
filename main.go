package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vx416/modelgen/pkg/modelgen"
	"github.com/vx416/modelgen/pkg/output"
	"github.com/vx416/modelgen/pkg/setting"
)

var s = &setting.Settings{}

func newSettings() {
	flag.StringVar(&s.InputPath, "i", "", "file path of schema file")
	flag.StringVar(&s.Destination, "o", "", "destination of output content")
	flag.StringVar(&s.Tag, "tag", "db", "model tag")
	flag.StringVar(&s.PackageName, "package", "model", "package of model")
	flag.StringVar(&s.DBKind, "db", "mysql", "kind of database")

	flag.BoolVar(&s.Print, "print", true, "print out of generated content")
	flag.BoolVar(&s.AppendOnly, "append", false, "append only")
	flag.Parse()
}

func main() {
	newSettings()

	data, err := ioutil.ReadFile(s.InputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	models, err := modelgen.NewModels(s, string(data))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := output.Output(s.AppendOnly, s.PackageName, models)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if s.Print {
		fmt.Println(res)
	}
	if s.Destination != "" {
		outputFile, err := os.Create(s.Destination)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outputFile.WriteString(res)
	}
}
