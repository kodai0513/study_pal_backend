package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/selection_problems"
	"study-pal-backend/app/usecases/shared/trancaction"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SelectionProblemController struct {
	AppData *app_types.AppData
}

type UpdatedAtSelectionProblem struct {
	SelectionProblemAnswers []*UpdateSelectionProblemAnswer `json:"selection_problem_answers"`
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
//	@Param		selection_problem_id	path		string							true	"SelectionProblem ID"
//	@Success	200						{object}	UpdateSelectionProblemResponse
//	@Failure	400						{object}	app_types.ErrorResponse
//	@Failure	401						{object}	app_types.ErrorResponse
//	@Failure	404						{object}	app_types.ErrorResponse
//	@Failure	500						{object}	app_types.ErrorResponse
//	@Router		/selection-problems/{selection_problem_id} [patch]
func (a *SelectionProblemController) Update(c *gin.Context) {
	selectionProblemId, err := uuid.Parse(c.Param("selection_problem_id"))
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
	var request UpdateSelectionProblemRequest
	err = c.BindJSON(&request)
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

	invalidUuidErrors := make([]string, 0)
	answerCommands := lo.FilterMap(request.SelectionProblemAnswers, func(answer *UpdateSelectionProblemAnswer, _ int) (*selection_problems.SelectionProblemAnswer, bool) {
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}

		return &selection_problems.SelectionProblemAnswer{
			IsCorrect: answer.IsCorrect,
			Statement: answer.Statement,
		}, true
	})

	if len(invalidUuidErrors) > 0 {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: invalidUuidErrors,
			},
		)
		c.Abort()
		return
	}
	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &selection_problems.UpdateAction{
		SelectionProblemRepository: repositories.NewSelectionProblemRepositoryImpl(tx, c),
		Tx:                         trancaction.NewTx(tx),
	}
	problemDto, usecaseErrGroup := action.Execute(
		&selection_problems.UpdateActionCommand{
			SelectionProblemAnswers: answerCommands,
			SelectionProblemId:      selectionProblemId,
			Statement:               request.Statement,
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

	answerResponse := lo.Map(problemDto.SelectionProblemAnswers, func(answer *selection_problems.SelectionProblemAnswerDto, _ int) *UpdateSelectionProblemAnswer {
		return &UpdateSelectionProblemAnswer{
			IsCorrect: answer.IsCorrect,
			Statement: answer.Statement,
		}
	})

	c.SecureJSON(
		http.StatusOK,
		&UpdateSelectionProblemResponse{
			&UpdatedAtSelectionProblem{
				SelectionProblemAnswers: answerResponse,
				Statement:               problemDto.Statement,
			},
		},
	)
}

// selection-problem godoc
//
//	@Summary	API
//	@Description
//	@Tags		selection-problem
//	@Accept		json
//	@Produce	json
//	@Param		selection_problem_id	path		string	true	"SelectionProblem ID"
//	@Success	204						{object}	nil
//	@Failure	400						{object}	app_types.ErrorResponse
//	@Failure	401						{object}	app_types.ErrorResponse
//	@Failure	404						{object}	app_types.ErrorResponse
//	@Failure	500						{object}	app_types.ErrorResponse
//	@Router		/selection-problems/{selection_problem_id} [delete]
func (a *SelectionProblemController) Delete(c *gin.Context) {
	selectionProblemId, err := uuid.Parse(c.Param("selection_problem_id"))
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
	action := &selection_problems.DeleteAction{
		SelectionProblemRepository: repositories.NewSelectionProblemRepositoryImpl(tx, c),
		Tx:                         trancaction.NewTx(tx),
	}
	usecaseErrGroup := action.Execute(
		&selection_problems.DeleteActionCommand{
			SelectionProblemId: selectionProblemId,
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
