package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/workbooks"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkbookController struct {
	appData *app_types.AppData
}

func NewWorkbookController(appData *app_types.AppData) *WorkbookController {
	return &WorkbookController{
		appData: appData,
	}
}

type WorkbookCreateRequest struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

// workbook godoc
//
//	@Summary	API
//	@Description
//	@Tags		workbook
//	@Accept		json
//	@Produce	json
//	@Param		request	body		WorkbookCreateRequest	true	"問題集作成リクエスト"
//	@Success	201		{object}	nil
//	@Failure	400		{object}	app_types.ErrorResponse
//	@Failure	401		{object}	app_types.ErrorResponse
//	@Failure	500		{object}	app_types.ErrorResponse
//	@Router		/workbooks [post]
func (a *WorkbookController) Create(c *gin.Context) {
	var request WorkbookCreateRequest
	c.BindJSON(&request)
	userId, _ := c.Get("user_id")
	usecaseErrGroup := workbooks.NewCreateAction(repositories.NewWorkbookRepositoryImpl(c, a.appData.Client())).Execute(
		workbooks.NewCreateActionCommand(request.Description, request.Title, userId.(uuid.UUID)),
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			app_types.NewErrorResponse(usecaseErrGroup.Errors()),
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusCreated,
		gin.H{},
	)
}
