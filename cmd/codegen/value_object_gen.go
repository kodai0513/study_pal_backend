package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
)

type fieldType string

const (
	String fieldType = "string"
	Int    fieldType = "int"
)

func getFieldType(value string) fieldType {
	switch value {
	case "string":
		return String
	case "int":
		return Int
	default:
		panic("invalid fieldType")
	}
}

func getErrValue(value fieldType) string {
	switch value {
	case String:
		return `""`
	case Int:
		return "0"
	default:
		panic("invalid fieldType")
	}
}

type valueObjectCommand struct {
	currentDir      string
	entityName      string
	valueObjectName string
	fieldType       fieldType
}

func valueObjectGen(command *valueObjectCommand) {
	valueObjectData := map[string]string{
		"errValue":        getErrValue(command.fieldType),
		"fieldType":       string(command.fieldType),
		"packageName":     camelToSnake(pluralize.NewClient().Plural(command.entityName)),
		"valueObjectName": command.valueObjectName,
	}

	templByte, err := os.ReadFile("templates/value_object.tmpl")
	if err != nil {
		panic(err)
	}

	outputPath := command.currentDir +
		"/../../app/domains/models/value_objects/" +
		camelToSnake(pluralize.NewClient().Plural(command.entityName)) +
		"/" +
		camelToSnake(command.valueObjectName) + ".go"

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

	err = tmpl.Execute(outputFile, valueObjectData)
	if err != nil {
		panic(err)
	}

	println("the value_object file has been successfully created")
}
