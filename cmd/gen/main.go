package main

import (
	"fmt"
	"os"
	"text/template"
)

// TemplateData is the input for the makefile template.
type TemplateData struct {
	ProjectType            string
	ProjectTypeDescription string
	LDFlags                bool
	GoVersion              bool
	Clean                  bool
	Build                  bool
	BuildPackage           string
	Test                   bool
	GoTestSum              bool
	Lint                   bool
	Release                bool
	TagPrefix              string
	KeepAChangeLog         bool
	TargetList             bool
	Help                   bool
	VarExport              bool
}

var templateData = []TemplateData{
	{
		ProjectType:            "go-app",
		ProjectTypeDescription: "a Go app project",
		LDFlags:                true,
		GoVersion:              true,
		Clean:                  true,
		Build:                  true,
		BuildPackage:           "./cmd/...",
		Test:                   true,
		GoTestSum:              true,
		Lint:                   true,
		Release:                true,
		TagPrefix:              "v",
		KeepAChangeLog:         true,
		TargetList:             true,
		Help:                   true,
		VarExport:              true,
	},
	{
		ProjectType:            "go-binary",
		ProjectTypeDescription: "a single-binary Go project",
		LDFlags:                true,
		GoVersion:              true,
		Clean:                  true,
		Build:                  true,
		BuildPackage:           ".",
		Test:                   true,
		GoTestSum:              true,
		Lint:                   true,
		Release:                true,
		TagPrefix:              "v",
		KeepAChangeLog:         true,
		TargetList:             true,
		Help:                   true,
		VarExport:              true,
	},
	{
		ProjectType:            "go-library",
		ProjectTypeDescription: "a Go library project",
		LDFlags:                false,
		Clean:                  false,
		Build:                  false,
		BuildPackage:           "", // noop
		Test:                   true,
		GoTestSum:              true,
		Lint:                   true,
		Release:                true,
		TagPrefix:              "v",
		KeepAChangeLog:         true,
		TargetList:             true,
		Help:                   true,
		VarExport:              true,
	},
}

func main() {
	tpl := template.Must(template.ParseFiles("template.mk"))

	for _, data := range templateData {
		file, err := os.OpenFile(fmt.Sprintf("../../%s/main.mk", data.ProjectType), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}

		err = tpl.Execute(file, data)
		if err != nil {
			panic(err)
		}
	}
}
