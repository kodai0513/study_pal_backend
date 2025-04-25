package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/permission_guard"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/description_problems"
	"study-pal-backend/app/usecases/shared/trancaction"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DescriptionProblemController struct {
	AppData *app_types.AppData
}

type UpdateDescriptionProblem struct {
	CorrentStatement string `json:"correct_statement"`
	Statement        string `json:"statement"`
}

type UpdateDescriptionProblemRequest struct {
	*UpdateDescriptionProblem
}

type UpdateDescriptionProblemResponse struct {
	*UpdateDescriptionProblem
}

// descriptionproblem godoc
//
//	@Summary		API
//	@Description	問題を編集します
//	@Tags			description-problem
//	@Accept			json
//	@Produce		json
//	@Param			request					body		UpdateDescriptionProblemRequest	true	"記述問題更新リクエスト"
//	@Param			workbook_id				path		string							true	"Workbook ID"
//	@Param			description_problem_id	path		string							true	"DescriptionProblem ID"
//	@Success		200						{object}	UpdateDescriptionProblemResponse
//	@Failure		400						{object}	app_types.ErrorResponse
//	@Failure		401						{object}	app_types.ErrorResponse
//	@Failure		403						{object}	app_types.ErrorResponse
//	@Failure		404						{object}	app_types.ErrorResponse
//	@Failure		500						{object}	app_types.ErrorResponse
//	@Router			/workbooks/{workbook_id}/description-problems/{description_problem_id} [patch]
func (a *DescriptionProblemController) Update(c *gin.Context) {
	var request UpdateDescriptionProblemRequest
	err := c.BindJSON(&request)
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

	workbookId, err := uuid.Parse(c.Param("workbook_id"))
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

	descriptionProblemId, err := uuid.Parse(c.Param("description_problem_id"))
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

	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &description_problems.UpdateAction{
		DescriptionProblemRepository: repositories.NewDescriptionProblemRepositoryImpl(tx, c),
		PermissionGuard:              permission_guard.NewWorkbookPermissionGuard(a.AppData.Client(), c),
		Tx:                           trancaction.NewTx(tx),
	}
	descriptionProblemDto, usecaseErrGroup := action.Execute(
		&description_problems.UpdateActionCommand{
			DescriptionProblemId: descriptionProblemId,
			CorrentStatement:     request.CorrentStatement,
			Statement:            request.Statement,
			UserId:               userId.(uuid.UUID),
			WorkbookId:           workbookId,
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
		&UpdateDescriptionProblemResponse{
			&UpdateDescriptionProblem{
				CorrentStatement: descriptionProblemDto.CorrentStatement,
				Statement:        descriptionProblemDto.Statement,
			},
		},
	)
}

// descriptionproblem godoc
//
//	@Summary	API
//	@Description
//	@Tags		description-problem
//	@Accept		json
//	@Produce	json
//	@Param		workbook_id				path		string	true	"Workbook ID"
//	@Param		description_problem_id	path		string	true	"DescriptionProblem ID"
//	@Success	204						{object}	nil
//	@Failure	400						{object}	app_types.ErrorResponse
//	@Failure	401						{object}	app_types.ErrorResponse
//	@Failure	403						{object}	app_types.ErrorResponse
//	@Failure	404						{object}	app_types.ErrorResponse
//	@Failure	500						{object}	app_types.ErrorResponse
//	@Router		/workbooks/{workbook_id}/description-problems/{description_problem_id} [delete]
func (a *DescriptionProblemController) Delete(c *gin.Context) {
	userId, _ := c.Get("user_id")

	workbookId, err := uuid.Parse(c.Param("workbook_id"))
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
	descriptionProblemId, err := uuid.Parse(c.Param("description_problem_id"))
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

	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &description_problems.DeleteAction{
		DescriptionProblemRepository: repositories.NewDescriptionProblemRepositoryImpl(tx, c),
		PermissionGuard:              permission_guard.NewWorkbookPermissionGuard(a.AppData.Client(), c),
		Tx:                           trancaction.NewTx(tx),
	}
	usecaseErrGroup := action.Execute(
		&description_problems.DeleteActionCommand{
			DescriptionProblemId: descriptionProblemId,
			UserId:               userId.(uuid.UUID),
			WorkbookId:           workbookId,
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
