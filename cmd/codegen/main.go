package main

import (
	"fmt"
	"os"
)

type actionType int

const (
	action actionType = iota
	controller
	entity
	repository
	query
	valueObject
)

func getCommandType(commandType string) actionType {
	switch commandType {
	case "action":
		return action
	case "controller":
		return controller
	case "entity":
		return entity
	case "query":
		return query
	case "repository":
		return repository
	case "valueObject":
		return valueObject
	default:
		panic("invalid command")
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("missing argument")
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if getCommandType(os.Args[1]) == action {
		actionGen(&actionCommand{
			currentDir: currentDir,
			name:       os.Args[2],
		})
	}

	if getCommandType(os.Args[1]) == controller {
		controllerGen(&controllerCommand{
			currentDir: currentDir,
			name:       os.Args[2],
		})
	}

	if getCommandType(os.Args[1]) == entity {
		entityGen(&entityCommand{
			currentDir: currentDir,
			entityName: os.Args[2],
			option:     os.Args[3],
		})
	}

	if getCommandType(os.Args[1]) == query {
		queryGen(&queryCommand{
			currentDir: currentDir,
			name:       os.Args[2],
		})
	}

	if getCommandType(os.Args[1]) == repository {
		repositoryGen(&repositoryCommand{
			currentDir: currentDir,
			name:       os.Args[2],
		})
	}

	if getCommandType(os.Args[1]) == valueObject {
		valueObjectGen(&valueObjectCommand{
			currentDir:      currentDir,
			entityName:      os.Args[2],
			valueObjectName: os.Args[3],
			fieldType:       getFieldType(os.Args[4]),
		})
	}
}
