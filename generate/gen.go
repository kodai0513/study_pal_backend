package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"slices"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
)

type generateInfo struct {
	commandType string
	name        string
}

func parseArgument(args []string) (*generateInfo, error) {
	if !slices.Contains([]string{"controller", "action", "query"}, strings.ToLower(args[1])) {
		return nil, errors.New("invalid argument")
	}
	return &generateInfo{
		commandType: strings.ToLower(args[1]),
		name:        args[2],
	}, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("missing argument")
		return
	}

	generateInfo, err := parseArgument(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if generateInfo.commandType == "action" {
		generateAction(generateInfo)
	}

	if generateInfo.commandType == "controller" {
		generateController(generateInfo)
	}

	if generateInfo.commandType == "query" {
		generateQuery(generateInfo)
	}
}

func generateAction(generateInfo *generateInfo) {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for index, actionType := range []string{"Create", "Update", "Delete"} {
		actionData := map[string]string{
			"actionCommandName": actionType + "ActionCommand",
			"actionName":        actionType + "Action",
			"packageName":       camelToSnake(pluralize.NewClient().Plural(generateInfo.name)),
			"dtoName":           generateInfo.name + "Dto",
		}

		var fileName string
		if actionType == "Create" || actionType == "Update" {
			fileName = "templates/create_or_update_action.tmpl"
		} else {
			fileName = "templates/delete_action.tmpl"
		}
		templByte, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		actionOutputPath := currentPath + "/../app/usecases/" + camelToSnake(pluralize.NewClient().Plural(generateInfo.name)) + "/" + strings.ToLower(actionType) + "_action.go"
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
		"packageName": camelToSnake(pluralize.NewClient().Plural(generateInfo.name)),
		"dtoName":     generateInfo.name + "Dto",
	}
	dtoOutputPath := currentPath + "/../app/usecases/" + camelToSnake(pluralize.NewClient().Plural(generateInfo.name)) + "/" + strings.ToLower(generateInfo.name) + "_action.go"

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

func generateController(generateInfo *generateInfo) {
	data := map[string]string{
		"controllerName":     generateInfo.name + "Controller",
		"godocName":          strings.ToLower(generateInfo.name),
		"createRequestName":  "Create" + generateInfo.name + "Request",
		"createResponseName": "Create" + generateInfo.name + "Response",
		"updateRequestName":  "Update" + generateInfo.name + "Request",
		"updateResponseName": "Update" + generateInfo.name + "Response",
	}

	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	templByte, err := os.ReadFile("templates/controller.tmpl")
	if err != nil {
		panic(err)
	}

	outputPath := currentPath + "/../app/controllers/" + camelToSnake(generateInfo.name) + "_controller.go"
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

func generateQuery(generateInfo *generateInfo) {
	data := map[string]string{
		"dtoName":          strings.ToUpper(generateInfo.name[0:1]) + strings.ToLower(generateInfo.name[1:]) + "Dto",
		"packageName":      camelToSnake(pluralize.NewClient().Plural(generateInfo.name)),
		"queryServiceName": strings.ToUpper(generateInfo.name[0:1]) + strings.ToLower(generateInfo.name[1:]) + "QueryService",
	}

	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	templByte, err := os.ReadFile("templates/query_service.tmpl")
	if err != nil {
		panic(err)
	}

	outputPath := currentPath + "/../app/usecases/" + camelToSnake(pluralize.NewClient().Plural(generateInfo.name)) + "/" + camelToSnake(generateInfo.name) + "_query_service.go"
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

	println("the query_service file has been successfully created")
}

func camelToSnake(s string) string {
	if s == "" {
		return s
	}

	delimiter := "_"
	sLen := len(s)
	var snake string
	for i, current := range s {
		if i > 0 && i+1 < sLen {
			if current >= 'A' && current <= 'Z' {
				next := s[i+1]
				prev := s[i-1]
				if (next >= 'a' && next <= 'z') || (prev >= 'a' && prev <= 'z') {
					snake += delimiter
				}
			}
		}
		snake += string(current)
	}

	snake = strings.ToLower(snake)
	return snake
}
