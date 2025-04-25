package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/permission_guard"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/true_or_false_problems"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TrueOrFalseProblemController struct {
	AppData *app_types.AppData
}

type UpdateTrueOrFalseProblem struct {
	IsCorrect bool   `json:"is_correct"`
	Statement string `json:"statement"`
}

type UpdateTrueOrFalseProblemRequest struct {
	*UpdateTrueOrFalseProblem
}

type UpdateTrueOrFalseProblemResponse struct {
	*UpdateTrueOrFalseProblem
}

// true-or-false-problem godoc
//
//	@Summary	API
//	@Description
//	@Tags		true-or-false-problem
//	@Accept		json
//	@Produce	json
//	@Param		request						body		UpdateTrueOrFalseProblem	true	"正誤問題更新リクエスト"
//	@Param		workbook_id					path		string						true	"Workbook ID"
//	@Param		true_or_false_problem_id	path		string						true	"TrueOrFalseProblem ID"
//	@Success	200							{object}	UpdateTrueOrFalseProblemResponse
//	@Failure	400							{object}	app_types.ErrorResponse
//	@Failure	401							{object}	app_types.ErrorResponse
//	@Failure	403							{object}	app_types.ErrorResponse
//	@Failure	404							{object}	app_types.ErrorResponse
//	@Failure	500							{object}	app_types.ErrorResponse
//	@Router		/workbooks/{workbook_id}/true-or-false-problems/{true_or_false_problem_id} [patch]
func (a *TrueOrFalseProblemController) Update(c *gin.Context) {
	var request UpdateTrueOrFalseProblemRequest
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
	}

	trueOrFalseProblemId, err := uuid.Parse(c.Param("true_or_false_problem_id"))
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
	action := &true_or_false_problems.UpdateAction{
		PermissionGuard:              permission_guard.NewWorkbookPermissionGuard(a.AppData.Client(), c),
		TrueOrFalseProblemRepository: repositories.NewTrueOrFalseProblemRepositoryImpl(tx, c),
		Tx:                           trancaction.NewTx(tx),
	}
	trueOrFalseProblemDto, usecaseErrGroup := action.Execute(
		&true_or_false_problems.UpdateActionCommand{
			IsCorrect:            request.IsCorrect,
			Statement:            request.Statement,
			TrueOrFalseProblemId: trueOrFalseProblemId,
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
		&UpdateTrueOrFalseProblemResponse{
			&UpdateTrueOrFalseProblem{
				IsCorrect: trueOrFalseProblemDto.IsCorrect,
				Statement: trueOrFalseProblemDto.Statement,
			},
		},
	)
}

// true-or-false-problem godoc
//
//	@Summary	API
//	@Description
//	@Tags		true-or-false-problem
//	@Accept		json
//	@Produce	json
//	@Param		workbook_id					path		string	true	"Workbook ID"
//	@Param		true_or_false_problem_id	path		string	true	"TrueOrFalseProblem ID"
//	@Success	204							{object}	nil
//	@Failure	400							{object}	app_types.ErrorResponse
//	@Failure	401							{object}	app_types.ErrorResponse
//	@Failure	404							{object}	app_types.ErrorResponse
//	@Failure	500							{object}	app_types.ErrorResponse
//	@Router		/workbooks/{workbook_id}/true-or-false-problems/{true_or_false_problem_id} [delete]
func (a *TrueOrFalseProblemController) Delete(c *gin.Context) {
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
	trueOrFalseProblemId, err := uuid.Parse(c.Param("true_or_false_problem_id"))
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
	action := &true_or_false_problems.DeleteAction{
		PermissionGuard:              permission_guard.NewWorkbookPermissionGuard(a.AppData.Client(), c),
		TrueOrFalseProblemRepository: repositories.NewTrueOrFalseProblemRepositoryImpl(tx, c),
		Tx:                           trancaction.NewTx(tx),
	}
	usecaseErrGroup := action.Execute(
		&true_or_false_problems.DeleteActionCommand{
			TrueOrFalseProblemId: trueOrFalseProblemId,
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
