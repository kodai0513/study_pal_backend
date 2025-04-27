package controllers

import (
	"study-pal-backend/app/app_types"

	"github.com/gin-gonic/gin"
)

type WorkbookMemberController struct {
	AppData *app_types.AppData
}

type IndexWorkbookMemberResponse struct {}

// workbook-member godoc
//
//	@Summary		API
//	@Description    
//	@Tags			workbook-member
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	IndexWorkbookMemberResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [get]
func (a *WorkbookMemberController) Index(c *gin.Context) {
}

type CreateWorkbookMemberRequest struct {}

type CreateWorkbookMemberResponse struct {}

// workbook-member godoc
//
//	@Summary		API
//	@Description    
//	@Tags			workbook-member
//	@Accept			json
//	@Produce		json
//	@Param			request		body	CreateWorkbookMemberRequest	true	""
//	@Success		201		{object}	CreateWorkbookMemberResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [post]
func (a *WorkbookMemberController) Create(c *gin.Context) {
}


type UpdateWorkbookMemberRequest struct {}

type UpdateWorkbookMemberResponse struct {}

// workbook-member godoc
//
//	@Summary		API
//	@Description    
//	@Tags			workbook-member
//	@Accept			json
//	@Produce		json
//	@Param			request		body	UpdateWorkbookMemberRequest	true	""
//	@Success		200		{object}	UpdateWorkbookMemberResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [put]
func (a *WorkbookMemberController) Update(c *gin.Context) {
}


// workbook-member godoc
//
//	@Summary		API
//	@Description    
//	@Tags			workbook-member
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [delete]
func (a *WorkbookMemberController) Delete(c *gin.Context) {
}