package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
)

type repoGenerateInfo struct {
	outputFilePath   string
	templateFilePath string
}

type repositoryCommand struct {
	currentDir string
	name       string
}

func repositoryGen(command *repositoryCommand) {
	data := map[string]string{
		"repositoryName": strcase.ToCamel(command.name),
	}

	generateInfos := []*repoGenerateInfo{
		{
			outputFilePath: command.currentDir +
				"/../../app/domains/repositories/" +
				strcase.ToSnake(command.name) +
				"_repository.go",
			templateFilePath: "templates/repository.tmpl",
		},
		{
			outputFilePath: command.currentDir +
				"/../../app/infrastructures/repositories/" +
				strcase.ToSnake(command.name) +
				"_repository_impl.go",
			templateFilePath: "templates/repository_impl.tmpl",
		},
	}

	for _, generateInfo := range generateInfos {
		tmplByte, err := os.ReadFile(generateInfo.templateFilePath)
		if err != nil {
			panic(err)
		}
		if _, err := os.Stat(generateInfo.outputFilePath); err == nil {
			fmt.Println("file has already been created")
			return
		}
		os.Mkdir(filepath.Dir(generateInfo.outputFilePath), 0755)

		outputFile, err := os.Create(generateInfo.outputFilePath)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		tmpl, err := template.New("").Parse(string(tmplByte))
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outputFile, data)
		if err != nil {
			panic(err)
		}

	}
	println("the repository file has been successfully created")
}
