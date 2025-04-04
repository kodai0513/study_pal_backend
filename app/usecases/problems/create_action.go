package problems

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/description_problems"
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"
	"study-pal-backend/app/domains/models/value_objects/true_or_false_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CreateDescriptionProblem struct {
	CorrentStatement         string
	Statement                string
	WorkbookCategoryDetailId *uuid.UUID
	WorkbookCategoryId       *uuid.UUID
}

type CreateSelectionProblem struct {
	SelectionProblemAnswers  []*CreateSelectionProblemAnswer
	Statement                string
	WorkbookCategoryDetailId *uuid.UUID
	WorkbookCategoryId       *uuid.UUID
}

type CreateSelectionProblemAnswer struct {
	IsCorrect bool
	Statement string
}

type CreateTrueOrFalseProblem struct {
	IsCorrect                bool
	Statement                string
	WorkbookCategoryDetailId *uuid.UUID
	WorkbookCategoryId       *uuid.UUID
}

type CreateActionCommand struct {
	DescriptionProblems []*CreateDescriptionProblem
	SelectionProblems   []*CreateSelectionProblem
	TrueOrFalseProblems []*CreateTrueOrFalseProblem
	WorkbookId          uuid.UUID
}

type CreateAction struct {
	DescriptionProblemRepository repositories.DescriptionProblemRepository
	SelectionProblemRepository   repositories.SelectionProblemRepository
	TrueOrFalseRepository        repositories.TrueOrFalseProblemRepository
	WorkbookRepository           repositories.WorkbookRepository
}

func (c *CreateAction) Execute(command *CreateActionCommand) (*ProblemDto, usecase_error.UsecaseErrorGroup) {
	hasWorkbook := c.WorkbookRepository.ExistById(command.WorkbookId)
	if !hasWorkbook {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbook not found")))
	}

	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	descriptionProblems := lo.Map(command.DescriptionProblems, func(problem *CreateDescriptionProblem, _ int) *entities.DescriptionProblem {
		correctStatement, err := description_problems.NewCorrectStatement(problem.CorrentStatement)
		if err != nil {
			usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}
		statement, err := description_problems.NewStatement(problem.Statement)
		if err != nil {
			usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}

		return entities.NewDescriptionProblem(
			uuid.New(),
			correctStatement,
			statement,
			problem.WorkbookCategoryDetailId,
			problem.WorkbookCategoryId,
			command.WorkbookId,
		)
	})
	selectoinProblems := lo.Map(command.SelectionProblems, func(problem *CreateSelectionProblem, _ int) *entities.SelectionProblem {
		statement, err := selection_problems.NewStatement(problem.Statement)
		if err != nil {
			usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}
		problemEntity := entities.CreateSelectionProblem(
			uuid.New(),
			statement,
			problem.WorkbookCategoryDetailId,
			problem.WorkbookCategoryId,
			command.WorkbookId,
		)
		answerEntities := lo.Map(problem.SelectionProblemAnswers, func(answer *CreateSelectionProblemAnswer, _ int) *entities.SelectionProblemAnswer {
			statement, err := selection_problem_answers.NewStatement(answer.Statement)
			if err != nil {
				usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
			}

			return entities.NewSelectionProblemAnswer(
				uuid.New(),
				answer.IsCorrect,
				problemEntity.Id(),
				statement,
			)
		})

		err = problemEntity.ReplaceSelectionProblemAnswer(answerEntities)
		if err != nil {
			usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}

		return problemEntity
	})
	trueOrFalseProblems := lo.Map(command.TrueOrFalseProblems, func(problem *CreateTrueOrFalseProblem, _ int) *entities.TrueOrFalseProblem {
		statement, err := true_or_false_problems.NewStatement(problem.Statement)
		if err != nil {
			usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}

		return entities.NewTrueOrFalseProblem(
			uuid.New(),
			problem.IsCorrect,
			statement,
			problem.WorkbookCategoryDetailId,
			problem.WorkbookCategoryId,
			command.WorkbookId,
		)
	})

	if usecaseErrGroup.IsError() {
		return nil, usecaseErrGroup
	}

	DescriptionProblemResults := c.DescriptionProblemRepository.CreateBulk(descriptionProblems)
	selectoinProblemResults := c.SelectionProblemRepository.CreateBulk(selectoinProblems)
	trueOrFalseProblemResults := c.TrueOrFalseRepository.CreateBulk(trueOrFalseProblems)
	descriptionProblemDtos := lo.Map(DescriptionProblemResults, func(problem *entities.DescriptionProblem, _ int) *DescriptionProblemDto {
		return &DescriptionProblemDto{
			CorrentStatement:         problem.CorrectStatement(),
			Statement:                problem.Statement(),
			WorkbookCategoryDetailId: problem.WorkbookCategoryId(),
			WorkbookCategoryId:       problem.WorkbookCategoryId(),
		}
	})
	selectionProblemDtos := lo.Map(selectoinProblemResults, func(problem *entities.SelectionProblem, _ int) *SelectionProblemDto {
		answerDtos := lo.Map(problem.SelectionProblemAnswers(), func(answer *entities.SelectionProblemAnswer, _ int) *SelectionProblemAnswerDto {
			return &SelectionProblemAnswerDto{
				IsCorrect: answer.IsCorrect(),
				Statement: answer.Statement(),
			}
		})

		return &SelectionProblemDto{
			SelectionProblemAnswers:  answerDtos,
			Statement:                problem.Statement(),
			WorkbookCategoryDetailId: problem.WorkbookCategoryDetailId(),
			WorkbookCategoryId:       problem.WorkbookCategoryId(),
		}
	})
	trueOrFalseProblemDtos := lo.Map(trueOrFalseProblemResults, func(problem *entities.TrueOrFalseProblem, _ int) *TrueOrFalseProblemDto {
		return &TrueOrFalseProblemDto{
			IsCorrect:                problem.IsCorrect(),
			Statement:                problem.Statement(),
			WorkbookCategoryDetailId: problem.WorkbookCategoryDetailId(),
			WorkbookCategoryId:       problem.WorkbookCategoryId(),
		}
	})

	return &ProblemDto{
		DescriptionProblemDtos: descriptionProblemDtos,
		SelectionProblemDtos:   selectionProblemDtos,
		TrueOrFalseProblemDtos: trueOrFalseProblemDtos,
	}, nil
}
