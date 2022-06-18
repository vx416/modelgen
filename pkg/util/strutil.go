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
		if strings.EqualFold(s, "id") {
			return "ID"
		}
		return caser.String(s)
	}

	var cc string
	for _, part := range splitted {
		if strings.EqualFold(part, "id") {
			cc += "ID"
			continue
		}
		cc += caser.String(strings.ToLower(part))
	}
	return cc
}

func LowercaseCamelCaseString(s string) string {
	if s == "" {
		return s
	}
	splitted := strings.Split(s, "_")

	if len(splitted) == 1 {
		return s
	}

	var cc string
	for i, part := range splitted {
		if strings.EqualFold(part, "id") {
			cc += "id"
			continue
		}
		if i == 0 {
			cc += strings.ToLower(part)
			continue
		}
		cc += caser.String(strings.ToLower(part))
	}
	return cc
}
