package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type tableColumn struct {
	name      string
	fieldType string
}

type valueObjectColumn struct {
	name      string
	fieldType string
}

type entityField struct {
	Name      string
	FieldType string
}

type entityGetter struct {
	FieldName        string
	GetterFuncName   string
	GetterReturnType string
	EntityName       string
	IsValueObject    bool
	VarEntityName    string
}

type entityCommand struct {
	currentDir string
	entityName string
	option     string
}

func entityGen(command *entityCommand) {
	valueObjectFiles, _ := os.ReadDir(command.currentDir + "/../../app/domains/models/value_objects/" + strcase.ToSnake(pluralize.NewClient().Plural(command.entityName)))

	valueObjectColumns := make(map[string]*valueObjectColumn, 0)
	for _, v := range valueObjectFiles {
		if !v.IsDir() {
			fileName := v.Name()
			ext := filepath.Ext(fileName)
			nameWithoutExt := strings.TrimSuffix(fileName, ext)
			valueObjectColumns[strcase.ToCamel(nameWithoutExt)] = &valueObjectColumn{
				name:      strcase.ToCamel(nameWithoutExt),
				fieldType: strcase.ToSnake(pluralize.NewClient().Plural(command.entityName)) + "." + strcase.ToCamel(nameWithoutExt),
			}
		}
	}

	entityFields := make([]*entityField, 0)
	entityGetters := make([]*entityGetter, 0)

	if command.option != "--no-field" {
		entFieldColumns := getEntField(command.entityName, command.currentDir+"/../../ent/"+strings.ToLower(command.entityName)+".go")
		if len(entFieldColumns) == 0 {
			panic("ent structure has not been created")
		}

		for _, entField := range entFieldColumns {
			valObjField, ok := valueObjectColumns[entField.name]
			if ok {
				entityFields = append(entityFields, &entityField{
					Name:      strcase.ToLowerCamel(valObjField.name),
					FieldType: valObjField.fieldType,
				})
				entityGetters = append(entityGetters, &entityGetter{
					FieldName:        strcase.ToLowerCamel(valObjField.name),
					GetterFuncName:   strcase.ToCamel(valObjField.name),
					GetterReturnType: entField.fieldType,
					IsValueObject:    true,
					VarEntityName:    strings.ToLower(command.entityName[0:1]),
					EntityName:       strcase.ToCamel(command.entityName),
				})
			} else {
				entityFields = append(entityFields, &entityField{
					Name:      strcase.ToLowerCamel(entField.name),
					FieldType: entField.fieldType,
				})
				entityGetters = append(entityGetters, &entityGetter{
					FieldName:        strcase.ToLowerCamel(entField.name),
					GetterFuncName:   strcase.ToCamel(entField.name),
					GetterReturnType: entField.fieldType,
					IsValueObject:    false,
					VarEntityName:    strings.ToLower(command.entityName[0:1]),
					EntityName:       strcase.ToCamel(command.entityName),
				})
			}
		}
	}

	data := map[string]any{
		"entityFields":     entityFields,
		"entityGetters":    entityGetters,
		"existValueObject": len(valueObjectColumns) != 0,
		"importEntityName": strcase.ToSnake(pluralize.NewClient().Plural(command.entityName)),
		"structEntityName": strcase.ToCamel(command.entityName),
	}

	templByte, err := os.ReadFile("templates/entity.tmpl")
	if err != nil {
		panic(err)
	}

	outputPath := command.currentDir +
		"/../../app/domains/models/entities/" +
		strcase.ToSnake(command.entityName) +
		".go"
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

	println("the entity file has been successfully created")
}
