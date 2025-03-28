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
	AppData *app_types.AppData
}

type CreateWorkbookRequest struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

type CreateWorkbookResponse struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	Title       string    `json:"title"`
	UserId      uuid.UUID `json:"user_id"`
}

// workbook godoc
//
//	@Summary		問題周作成API
//	@Description	問題集を作成します
//	@Tags			workbook
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateWorkbookRequest	true	"問題集作成リクエスト"
//	@Success		201		{object}	CreateWorkbookResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/workbooks [post]
func (a *WorkbookController) Create(c *gin.Context) {
	var request CreateWorkbookRequest
	c.BindJSON(&request)
	userId, _ := c.Get("user_id")
	action := workbooks.CreateAction{
		WorkbookRepository: repositories.NewWorkbookRepositoryImpl(a.AppData.Client(), c),
	}
	workbookDto, usecaseErrGroup := action.Execute(
		&workbooks.CreateActionCommand{
			Description: request.Description,
			Title:       request.Title,
			UserId:      userId.(uuid.UUID),
		},
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusCreated,
		&CreateWorkbookResponse{
			Id:          workbookDto.Id,
			Description: workbookDto.Description,
			IsPublic:    workbookDto.IsPublic,
			Title:       workbookDto.Title,
			UserId:      workbookDto.UserId,
		},
	)
}

type UpdateWorkbookRequest struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}

type UpdateWorkbookResponse struct {
	Id          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	Title       string    `json:"title"`
	UserId      uuid.UUID `json:"user_id"`
}

// workbook godoc
//
//	@Summary		問題集編集API
//	@Description	問題集を編集します
//	@Tags			workbook
//	@Accept			json
//	@Produce		json
//	@Param			request		body		UpdateWorkbookRequest	true	"問題集編集リクエスト"
//	@Param			workbook_id	path		string					true	"Workbook ID"
//	@Success		200			{object}	UpdateWorkbookResponse
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/workbooks/{workbook_id} [put]
func (a *WorkbookController) Update(c *gin.Context) {
	var request UpdateWorkbookRequest
	c.BindJSON(&request)
	workbookIdParam := c.Param("workbook_id")
	workbookId, err := uuid.Parse(workbookIdParam)
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}
	userId, _ := c.Get("user_id")
	action := workbooks.UpdateAction{
		WorkbookRepository: repositories.NewWorkbookRepositoryImpl(a.AppData.Client(), c),
	}
	workbookDto, usecaseErrGroup := action.Execute(
		&workbooks.UpdateActionCommand{
			Description: request.Description,
			Title:       request.Title,
			UserId:      userId.(uuid.UUID),
			WorkbookId:  workbookId,
		},
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusOK,
		&UpdateWorkbookResponse{
			Id:          workbookDto.Id,
			Description: workbookDto.Description,
			IsPublic:    workbookDto.IsPublic,
			Title:       workbookDto.Title,
			UserId:      workbookDto.UserId,
		},
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
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}
	userId, _ := c.Get("user_id")
	action := workbooks.DeleteAction{
		WorkbookRepository: repositories.NewWorkbookRepositoryImpl(a.AppData.Client(), c),
	}
	usecaseErrGroup := action.Execute(
		&workbooks.DeleteActionCommand{
			UserId: userId.(uuid.UUID),

			WorkbookId: workbookId,
		},
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusNoContent,
		gin.H{},
	)
}
