package main

import (
	"bytes"
	_ "embed"
	"os"
	"text/template"
)

//go:embed default.tpl
var errorsTemplate string

type errorInfo struct {
	Name       string
	Value      string
	Code       int
	HTTPCode   int
	Message    string
	CamelValue string
	Comment    string
	HasComment bool
}

type errorWrapper struct {
	Errors []*errorInfo
}

func InitErrorsTemplate(tmlPath string) error {
	if tmlPath != "" {
		b, err := os.ReadFile(tmlPath)
		if err != nil {
			return err
		}
		errorsTemplate = string(b)
	}
	return nil
}

func (e *errorWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errorsTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
