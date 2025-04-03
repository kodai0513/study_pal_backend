package controllers

import (
	"study-pal-backend/app/app_types"

	"github.com/gin-gonic/gin"
)

type SelectionProblemController struct {
	AppData *app_types.AppData
}

type UpdatedAtSelectionProblem struct {
	SelectionProblemAnswers []*CreateSelectionProblemAnswer `json:"selection_problem_answers"`
	Statement               string                          `json:"statement"`
}

type UpdateSelectionProblemAnswer struct {
	IsCorrect bool   `json:"is_correct"`
	Statement string `json:"statement"`
}

type UpdateSelectionProblemRequest struct {
	*UpdatedAtSelectionProblem
}

type UpdateSelectionProblemResponse struct {
	*UpdatedAtSelectionProblem
}

// selection-problem godoc
//
//	@Summary	API
//	@Description
//	@Tags		selection-problem
//	@Accept		json
//	@Produce	json
//	@Param		request					body		UpdateSelectionProblemRequest	true	"記述問題更新リクエスト"
//	@Param		description_problem_id	path		string							true	"SelectionProblem ID"
//	@Success	200						{object}	UpdateSelectionProblemResponse
//	@Failure	400						{object}	app_types.ErrorResponse
//	@Failure	401						{object}	app_types.ErrorResponse
//	@Failure	500						{object}	app_types.ErrorResponse
//	@Router		/selection-problems/{selection_problem_id} [put]
func (a *SelectionProblemController) Update(c *gin.Context) {

}

// selection-problem godoc
//
//	@Summary	API
//	@Description
//	@Tags		selection-problem
//	@Accept		json
//	@Produce	json
//	@Param		description_problem_id	path		string	true	"SelectionProblem ID"
//	@Success	204						{object}	nil
//	@Failure	400						{object}	app_types.ErrorResponse
//	@Failure	401						{object}	app_types.ErrorResponse
//	@Failure	500						{object}	app_types.ErrorResponse
//	@Router		/selection-problems/{selection_problem_id} [delete]
func (a *SelectionProblemController) Delete(c *gin.Context) {
}
