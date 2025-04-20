package main

import (
	"fmt"
	"os"
	"text/template"
)

type controllerCommand struct {
	currentDir string
	name       string
}

func controllerGen(command *controllerCommand) {
	data := map[string]string{
		"controllerName":     command.name + "Controller",
		"godocName":          camelToKebab(command.name),
		"indexResponseName":  "Index" + command.name + "Response",
		"createRequestName":  "Create" + command.name + "Request",
		"createResponseName": "Create" + command.name + "Response",
		"updateRequestName":  "Update" + command.name + "Request",
		"updateResponseName": "Update" + command.name + "Response",
	}

	templByte, err := os.ReadFile("templates/controller.tmpl")
	if err != nil {
		panic(err)
	}

	outputPath := command.currentDir + "/../../app/controllers/" + camelToSnake(command.name) + "_controller.go"
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Println("file has already been created")
		return
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

	println("the controller file has been successfully created")
}
