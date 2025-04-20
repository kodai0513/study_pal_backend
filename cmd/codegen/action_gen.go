package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
)

type actionCommand struct {
	currentDir string
	name       string
}

func actionGen(command *actionCommand) {
	for index, actionType := range []string{"Index", "Create", "Update", "Delete"} {
		actionData := map[string]string{
			"actionCommandName": actionType + "ActionCommand",
			"actionName":        actionType + "Action",
			"packageName":       camelToSnake(pluralize.NewClient().Plural(command.name)),
			"dtoName":           command.name + "Dto",
		}

		var fileName string
		if actionType == "Index" || actionType == "Create" || actionType == "Update" {
			fileName = "templates/index_or_create_or_update_action.tmpl"
		} else {
			fileName = "templates/delete_action.tmpl"
		}
		templByte, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		actionOutputPath := command.currentDir + "/../../app/usecases/" + camelToSnake(pluralize.NewClient().Plural(command.name)) + "/" + strings.ToLower(actionType) + "_action.go"
		if _, err := os.Stat(actionOutputPath); err == nil {
			fmt.Println("file has already been created")
			return
		}

		if index == 0 {
			err = os.Mkdir(filepath.Dir(actionOutputPath), 0755)
			if err != nil {
				panic(err)
			}
		}

		outputFile, err := os.Create(actionOutputPath)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		tmpl, err := template.New("").Parse(string(templByte))
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outputFile, actionData)
		if err != nil {
			panic(err)
		}
	}

	dtoData := map[string]string{
		"packageName": camelToSnake(pluralize.NewClient().Plural(command.name)),
		"dtoName":     command.name + "Dto",
	}
	dtoOutputPath := command.currentDir + "/../../app/usecases/" + camelToSnake(pluralize.NewClient().Plural(command.name)) + "/" + camelToSnake(command.name) + "_dto.go"

	outputFile, err := os.Create(dtoOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	templByte, err := os.ReadFile("templates/dto.tmpl")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("").Parse(string(templByte))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(outputFile, dtoData)
	if err != nil {
		panic(err)
	}
	println("the action file has been successfully created")
}
