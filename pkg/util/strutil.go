package util

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	caser = cases.Title(language.English, cases.NoLower)
)

func CamelCaseString(s string) string {
	if s == "" {
		return s
	}
	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		if s == "id" {
			return "ID"
		}
		return caser.String(s)
	}

	var cc string
	for _, part := range splitted {
		if part == "id" {
			cc += "ID"
			continue
		}
		cc += caser.String(strings.ToLower(part))
	}
	return cc
}
