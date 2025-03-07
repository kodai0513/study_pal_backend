package main

import (
	"fmt"
	"net/http"
	"os"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers"
	"study-pal-backend/app/infrastructures/db"
	"study-pal-backend/app/middlewares"
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

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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

	jwtSecretKey := os.Getenv("JWT_SERCRET_KEY")
	appData := app_types.NewAppData(client, jwtSecretKey)

	authController := controllers.NewAuthController(appData)
	timelineController := controllers.NewTimelineController(appData)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authController.Login)
		v1.POST("/refresh-token", authController.RefreshToken)
		v1.GET("/timelines", timelineController.Index)
		v1.Use(middlewares.AuthRequired(appData.JwtSecretKey())).GET("/auth-test", func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{"message": "hello world"},
			)
		})
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
