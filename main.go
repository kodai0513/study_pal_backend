package main

import (
	"fmt"
	"os"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers"
	"study-pal-backend/app/infrastructures/db"
	_ "study-pal-backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			StudyPal API
//	@version		1.0
//	@description	StudyPalAPIサーバー
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	JWTAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
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
	v1 := r.Group("/api/v1")
	{
		v1.GET("/timelines", timelineController.Index)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
