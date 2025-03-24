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
//	@Summary		問題周作成API
//	@Description	問題集を作成します
//	@Tags			workbook
//	@Accept			json
//	@Produce		json
//	@Param			request	body		WorkbookCreateRequest	true	"問題集作成リクエスト"
//	@Success		201		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/workbooks [post]
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

type WorkbookUpdateRequest struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

// workbook godoc
//
//	@Summary		問題集編集API
//	@Description	問題集を編集します
//	@Tags			workbook
//	@Accept			json
//	@Produce		json
//	@Param			request		body		WorkbookCreateRequest	true	"問題集編集リクエスト"
//	@Param			workbook_id	path		string					true	"Workbook ID"
//	@Success		200			{object}	nil
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/workbooks/{workbook_id} [put]
func (a *WorkbookController) Update(c *gin.Context) {
	var request WorkbookUpdateRequest
	c.BindJSON(&request)
	workbookIdParam := c.Param("workbook_id")
	workbookId, err := uuid.Parse(workbookIdParam)
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
		c.Abort()
		return
	}
	userId, _ := c.Get("user_id")
	usecaseErrGroup := workbooks.NewUpdateAction(repositories.NewWorkbookRepositoryImpl(c, a.appData.Client())).Execute(
		workbooks.NewUpdateActionCommand(request.Description, request.Title, userId.(uuid.UUID), workbookId),
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
		http.StatusOK,
		gin.H{},
	)
}

// workbook godoc
//
//	@Summary		問題集削除API
//	@Description	問題集を削除します
//	@Tags			workbook
//	@Accept			json
//	@Produce		json
//	@Param			workbook_id	path		string	true	"Workbook ID"
//	@Success		204			{object}	nil
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/workbooks/{workbook_id} [delete]
func (a *WorkbookController) Delete(c *gin.Context) {
	workbookIdParam := c.Param("workbook_id")
	workbookId, err := uuid.Parse(workbookIdParam)
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
		c.Abort()
		return
	}
	userId, _ := c.Get("user_id")
	usecaseErrGroup := workbooks.NewDeleteAction(repositories.NewWorkbookRepositoryImpl(c, a.appData.Client())).Execute(
		workbooks.NewDeleteActionCommand(userId.(uuid.UUID), workbookId),
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
		http.StatusNoContent,
		gin.H{},
	)
}
