package controllers

import (
	"study-pal-backend/app/app_types"

	"github.com/gin-gonic/gin"
)

type WorkbookInvitationMemberController struct {
	AppData *app_types.AppData
}

type IndexWorkbookInvitationMemberResponse struct{}

// workbook-invitation-member godoc
//
//	@Summary		API
//	@Description
//	@Tags			workbook-invitation-member
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	IndexWorkbookInvitationMemberResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		403		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [get]
func (a *WorkbookInvitationMemberController) Index(c *gin.Context) {
}

type CreateWorkbookInvitationMemberRequest struct {
	RoleId     string
	UserId     string
	WorkbookId string
}

// workbook-invitation-member godoc
//
//	@Summary		API
//	@Description
//	@Tags			workbook-invitation-member
//	@Accept			json
//	@Produce		json
//	@Param			request		body	CreateWorkbookInvitationMemberRequest	true	""
//	@Success		201		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		403		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/workbooks/{workbook_id}/workbook-invitation-members [post]
func (a *WorkbookInvitationMemberController) Create(c *gin.Context) {
}

type UpdateWorkbookInvitationMemberRequest struct{}

// workbook-invitation-member godoc
//
//	@Summary		API
//	@Description
//	@Tags			workbook-invitation-member
//	@Accept			json
//	@Produce		json
//	@Param			request		body	UpdateWorkbookInvitationMemberRequest	true	""
//	@Success		200		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		403		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [put]
func (a *WorkbookInvitationMemberController) Update(c *gin.Context) {
}

// workbook-invitation-member godoc
//
//	@Summary		API
//	@Description
//	@Tags			workbook-invitation-member
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		403		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [delete]
func (a *WorkbookInvitationMemberController) Delete(c *gin.Context) {
}
