package main

import (
	"fmt"
	"os"
	"strconv"
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

	devModeLogin, _ := strconv.ParseBool(os.Getenv("DEV_MODE_LOGIN"))
	jwtSecretKey := os.Getenv("JWT_SERCRET_KEY")
	appData := app_types.NewAppData(client, jwtSecretKey)
	articleController := controllers.NewArticleController(appData)
	authController := controllers.NewAuthController(appData)
	timelineController := controllers.NewTimelineController(appData)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authRequired := middlewares.AuthRequired(devModeLogin, appData.JwtSecretKey())

	v1 := r.Group("/api/v1")
	{
		article := v1.Group("/articles")
		{
			article.POST("/", authRequired, articleController.Create)
			article.PUT("/:article_id", authRequired, articleController.Update)
			article.DELETE("/:article_id", authRequired, articleController.Delete)
		}
		v1.POST("/login", authController.Login)
		v1.POST("/refresh-token", authController.RefreshToken)
		v1.GET("/timelines", timelineController.Index)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
