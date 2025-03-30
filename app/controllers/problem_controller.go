package controllers

import (
	"study-pal-backend/app/app_types"

	"github.com/gin-gonic/gin"
)

type ProblemController struct {
	appData *app_types.AppData
}

type CreateProblemRequest struct {}

type CreateProblemResponse struct {}

// problem godoc
//
//	@Summary		API
//	@Description    
//	@Tags			problem
//	@Accept			json
//	@Produce		json
//	@Success		201		{object}	CreateProblemResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [post]
func (a *ProblemController) Create(c *gin.Context) {
}


type UpdateProblemRequest struct {}

type UpdateProblemResponse struct {}

// problem godoc
//
//	@Summary		API
//	@Description    
//	@Tags			problem
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	UpdateProblemResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [put]
func (a *ProblemController) Update(c *gin.Context) {
}


// problem godoc
//
//	@Summary		API
//	@Description    
//	@Tags			problem
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/ [delete]
func (a *ProblemController) Delete(c *gin.Context) {
}