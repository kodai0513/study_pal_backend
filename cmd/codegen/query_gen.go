package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type queryCommand struct {
	currentDir string
	name       string
}

func queryGen(command *queryCommand) {
	// query_service
	{
		data := map[string]string{
			"dtoName":          strcase.ToCamel(command.name) + "Dto",
			"packageName":      camelToSnake(pluralize.NewClient().Plural(command.name)),
			"queryServiceName": strcase.ToCamel(command.name) + "QueryService",
		}

		templByte, err := os.ReadFile("templates/query_service.tmpl")
		if err != nil {
			panic(err)
		}

		outputPath := command.currentDir +
			"/../../app/usecases/" +
			camelToSnake(pluralize.NewClient().Plural(command.name)) +
			"/" +
			camelToSnake(command.name) +
			"_query_service.go"
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Println("file has already been created")
			return
		}

		err = os.Mkdir(filepath.Dir(outputPath), 0755)
		if err != nil {
			panic(err)
		}

		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		tmpl, err := template.New("").Parse(string(templByte))
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outputFile, data)
		if err != nil {
			panic(err)
		}
	}

	// query_service_impl
	{
		data := map[string]string{
			"importQueryServiceName": strcase.ToSnake(pluralize.NewClient().Plural(command.name)),
			"queryServiceName":       strcase.ToCamel(command.name),
		}

		templByte, err := os.ReadFile("templates/query_service_impl.tmpl")
		if err != nil {
			panic(err)
		}

		outputPath := command.currentDir +
			"/../../app/infrastructures/query_services/" +
			strcase.ToSnake(command.name) +
			"_query_service.go"
		if _, err := os.Stat(outputPath); err == nil {
			fmt.Println("file has already been created")
			return
		}

		os.Mkdir(filepath.Dir(outputPath), 0755)

		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		tmpl, err := template.New("").Parse(string(templByte))
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outputFile, data)
		if err != nil {
			panic(err)
		}
	}

	println("the query_service file has been successfully created")
}
