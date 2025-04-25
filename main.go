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
	articleController := controllers.ArticleController{AppData: appData}
	articleLikeController := controllers.ArticleLikeController{AppData: appData}
	authController := controllers.AuthController{AppData: appData}
	descriptionProblemController := controllers.DescriptionProblemController{AppData: appData}
	problemController := controllers.ProblemController{AppData: appData}
	selectionProblemController := controllers.SelectionProblemController{AppData: appData}
	timelineController := controllers.TimelineController{AppData: appData}
	trueOrFalseProblemController := controllers.TrueOrFalseProblemController{AppData: appData}
	workbookController := controllers.WorkbookController{AppData: appData}
	workbookCategoryController := controllers.WorkbookCategoryController{AppData: appData}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authRequired := middlewares.AuthRequired(devModeLogin, appData.JwtSecretKey())

	v1 := r.Group("/api/v1")
	{
		articles := v1.Group("/articles")
		{
			articles.POST("/", authRequired, articleController.Create)
			articles.PATCH("/:article_id", authRequired, articleController.Update)
			articles.DELETE("/:article_id", authRequired, articleController.Delete)

			likes := articles.Group("/:article_id/likes")
			{
				likes.POST("", authRequired, articleLikeController.Create)
				likes.DELETE("/:article_like_id", authRequired, articleLikeController.Delete)
			}
		}
		v1.POST("/login", authController.Login)
		v1.POST("/refresh-token", authController.RefreshToken)
		v1.GET("/timelines", timelineController.Index)
		workbooks := v1.Group("/workbooks")
		{
			workbooks.POST("/", authRequired, workbookController.Create)
			workbooks.PATCH("/:workbook_id", authRequired, workbookController.Update)
			workbooks.DELETE("/:workbook_id", authRequired, workbookController.Delete)

			workbookGroups := workbooks.Group("/:workbook_id")
			{
				descriptionProblems := workbookGroups.Group("/description-problems")
				{
					descriptionProblems.PATCH("/:description_problem_id", authRequired, descriptionProblemController.Update)
					descriptionProblems.DELETE("/:description_problem_id", authRequired, descriptionProblemController.Delete)
				}
				selectionProblems := workbookGroups.Group("/selection-problems")
				{
					selectionProblems.PATCH("/:selection_problem_id", authRequired, selectionProblemController.Update)
					selectionProblems.DELETE("/:selection_problem_id", authRequired, selectionProblemController.Delete)
				}
				trueOrFalseProblems := workbookGroups.Group("/true-or-false-problems")
				{
					trueOrFalseProblems.PATCH("/:true_or_false_problem_id", authRequired, trueOrFalseProblemController.Update)
					trueOrFalseProblems.DELETE("/:true_or_false_problem_id", authRequired, trueOrFalseProblemController.Delete)
				}
				workbookGroups.POST("/problems", authRequired, problemController.Create)
				workbookCategories := workbookGroups.Group("/workbook-categories")
				{
					workbookCategories.GET("", authRequired, workbookCategoryController.Index)
					workbookCategories.PUT("", authRequired, workbookCategoryController.Update)
				}
			}
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
