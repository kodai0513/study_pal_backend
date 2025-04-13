package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/problems"
	"study-pal-backend/app/utils/type_converts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ProblemController struct {
	AppData *app_types.AppData
}

type CreateDescriptionProblem struct {
	CorrentStatement   string `json:"correct_statement"`
	Statement          string `json:"statement"`
	WorkbookCategoryId string `json:"workbook_category_id"`
}

type CreateSelectionProblem struct {
	SelectionProblemAnswers []*CreateSelectionProblemAnswer `json:"selection_problem_answers"`
	Statement               string                          `json:"statement"`
	WorkbookCategoryId      string                          `json:"workbook_category_id"`
}

type CreateSelectionProblemAnswer struct {
	IsCorrect bool   `json:"is_correct"`
	Statement string `json:"statement"`
}

type CreateTrueOrFalseProblem struct {
	IsCorrect          bool   `json:"is_correct"`
	Statement          string `json:"statement"`
	WorkbookCategoryId string `json:"workbook_category_id"`
}

type CreateProblem struct {
	DescriptionProblems []*CreateDescriptionProblem `json:"description_problems"`
	SelectionProblems   []*CreateSelectionProblem   `json:"selection_problems"`
	TrueOrFalseProblems []*CreateTrueOrFalseProblem `json:"true_or_false_problems"`
}

type CreateProblemRequest struct {
	*CreateProblem
}

type CreateProblemResponse struct {
	*CreateProblem
}

// problem godoc
//
//	@Summary		API
//	@Description	問題を作成します
//	@Tags			problem
//	@Accept			json
//	@Produce		json
//	@Param			request		body		CreateProblemRequest	true	"問題作成リクエスト"
//	@Param			workbook_id	path		string					true	"Workbook ID"
//	@Success		201			{object}	CreateProblemResponse
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		404			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/{workbook_id}/problems [post]
func (a *ProblemController) Create(c *gin.Context) {
	var request CreateProblemRequest
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
	//userId, _ := c.Get("user_id")
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

	invalidUuidErrors := make([]string, 0)
	descriptionProblemCommands := lo.Map(request.DescriptionProblems, func(problem *CreateDescriptionProblem, _ int) *problems.CreateDescriptionProblem {
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		workbookCategoryUuid, err := type_converts.StringToUuidOrNil(problem.WorkbookCategoryId)
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		return &problems.CreateDescriptionProblem{
			CorrentStatement:   problem.CorrentStatement,
			Statement:          problem.Statement,
			WorkbookCategoryId: workbookCategoryUuid,
		}
	})
	selectionProblemCommands := lo.Map(request.SelectionProblems, func(problem *CreateSelectionProblem, _ int) *problems.CreateSelectionProblem {
		answers := lo.Map(problem.SelectionProblemAnswers, func(answer *CreateSelectionProblemAnswer, _ int) *problems.CreateSelectionProblemAnswer {
			return &problems.CreateSelectionProblemAnswer{
				IsCorrect: answer.IsCorrect,
				Statement: answer.Statement,
			}
		})
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		workbookCategoryUuid, err := type_converts.StringToUuidOrNil(problem.WorkbookCategoryId)
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		return &problems.CreateSelectionProblem{
			SelectionProblemAnswers: answers,
			Statement:               problem.Statement,
			WorkbookCategoryId:      workbookCategoryUuid,
		}
	})
	trueOrFalseProblemCommands := lo.Map(request.TrueOrFalseProblems, func(problem *CreateTrueOrFalseProblem, _ int) *problems.CreateTrueOrFalseProblem {
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		workbookCategoryUuid, err := type_converts.StringToUuidOrNil(problem.WorkbookCategoryId)
		if err != nil {
			invalidUuidErrors = append(invalidUuidErrors, err.Error())
		}
		return &problems.CreateTrueOrFalseProblem{
			IsCorrect:          problem.IsCorrect,
			Statement:          problem.Statement,
			WorkbookCategoryId: workbookCategoryUuid,
		}
	})

	if len(invalidUuidErrors) != 0 {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: invalidUuidErrors,
			},
		)
		c.Abort()
		return
	}

	action := problems.CreateAction{
		DescriptionProblemRepository: repositories.NewDescriptionProblemRepositoryImpl(a.AppData.Client(), c),
		SelectionProblemRepository:   repositories.NewSelectionProblemRepositoryImpl(a.AppData.Client(), c),
		TrueOrFalseRepository:        repositories.NewTrueOrFalseProblemRepositoryImpl(a.AppData.Client(), c),
		WorkbookRepository:           repositories.NewWorkbookRepositoryImpl(a.AppData.Client(), c),
	}
	problemDto, usecaseErrGroup := action.Execute(
		&problems.CreateActionCommand{
			DescriptionProblems: descriptionProblemCommands,
			SelectionProblems:   selectionProblemCommands,
			TrueOrFalseProblems: trueOrFalseProblemCommands,
			WorkbookId:          workbookId,
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

	descriptionProblemReses := lo.Map(problemDto.DescriptionProblemDtos, func(problem *problems.DescriptionProblemDto, _ int) *CreateDescriptionProblem {
		return &CreateDescriptionProblem{
			CorrentStatement:   problem.CorrentStatement,
			Statement:          problem.Statement,
			WorkbookCategoryId: type_converts.PointerUuidToString(problem.WorkbookCategoryId),
		}
	})
	selectionProblemReses := lo.Map(problemDto.SelectionProblemDtos, func(problem *problems.SelectionProblemDto, _ int) *CreateSelectionProblem {
		answers := lo.Map(problem.SelectionProblemAnswers, func(answer *problems.SelectionProblemAnswerDto, _ int) *CreateSelectionProblemAnswer {
			return &CreateSelectionProblemAnswer{
				IsCorrect: answer.IsCorrect,
				Statement: answer.Statement,
			}
		})
		return &CreateSelectionProblem{
			SelectionProblemAnswers: answers,
			Statement:               problem.Statement,
			WorkbookCategoryId:      type_converts.PointerUuidToString(problem.WorkbookCategoryId),
		}
	})
	trueOrFalseDescriptionReses := lo.Map(problemDto.TrueOrFalseProblemDtos, func(problem *problems.TrueOrFalseProblemDto, _ int) *CreateTrueOrFalseProblem {
		return &CreateTrueOrFalseProblem{
			IsCorrect:          problem.IsCorrect,
			Statement:          problem.Statement,
			WorkbookCategoryId: type_converts.PointerUuidToString(problem.WorkbookCategoryId),
		}
	})

	c.SecureJSON(
		http.StatusCreated,
		&CreateProblemResponse{
			&CreateProblem{
				DescriptionProblems: descriptionProblemReses,
				SelectionProblems:   selectionProblemReses,
				TrueOrFalseProblems: trueOrFalseDescriptionReses,
			},
		},
	)
}
