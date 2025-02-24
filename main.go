package main

import (
	"fmt"
	"os"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers"
	"study-pal-backend/app/infrastructures/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		println(err.Error())
	}
	dbUser := os.Getenv("DB_USER")
	dbPassowrd := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databaseUrl := fmt.Sprintf("postgres:///%s?host=%s&port=%s&user=%s&password=%s", dbName, dbHost, dbPort, dbUser, dbPassowrd)

	client, err := db.Open(databaseUrl)
	if err != nil {
		println(err.Error())
		return
	}
	defer client.Close()

	appData := app_types.NewAppData(client)

	timelineController := controllers.NewTimelineController(appData)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	v1 := r.Group("/v1")
	{
		v1.GET("/timelines", timelineController.Index)
	}
	r.Run()
}
