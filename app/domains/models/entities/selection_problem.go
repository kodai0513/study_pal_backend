package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SelectionProblem struct {
	id                      uuid.UUID
	selectionProblemAnswers []*SelectionProblemAnswer
	statement               selection_problems.Statement
	workbookCategoryId      *uuid.UUID
	workbookId              uuid.UUID
}

func CreateSelectionProblem(
	id uuid.UUID,
	statement selection_problems.Statement,
	workbookCategoryId *uuid.UUID,
	workbookId uuid.UUID,
) *SelectionProblem {
	return &SelectionProblem{
		id:                      id,
		selectionProblemAnswers: make([]*SelectionProblemAnswer, 0),
		statement:               statement,
		workbookCategoryId:      workbookCategoryId,
		workbookId:              workbookId,
	}
}

func NewSelectionProblem(
	id uuid.UUID,
	selectionProblemAnswers []*SelectionProblemAnswer,
	statement selection_problems.Statement,
	workbookCategoryId *uuid.UUID,
	workbookId uuid.UUID,
) *SelectionProblem {
	return &SelectionProblem{
		id:                      id,
		selectionProblemAnswers: selectionProblemAnswers,
		statement:               statement,
		workbookCategoryId:      workbookCategoryId,
		workbookId:              workbookId,
	}
}

func (s *SelectionProblem) Id() uuid.UUID {
	return s.id
}

func (s *SelectionProblem) SelectionProblemAnswers() []*SelectionProblemAnswer {
	return s.selectionProblemAnswers
}

func (s *SelectionProblem) Statement() string {
	return s.statement.Value()
}

func (s *SelectionProblem) WorkbookCategoryId() *uuid.UUID {
	return s.workbookCategoryId
}

func (s *SelectionProblem) WorkbookId() uuid.UUID {
	return s.workbookId
}

func (s *SelectionProblem) ReplaceSelectionProblemAnswer(selectionProblemAnswers []*SelectionProblemAnswer) error {
	if len(selectionProblemAnswers) < 2 {
		return errors.New("at least 2 selectionProblemAnswer are required")
	}

	if len(selectionProblemAnswers) > 30 {
		return errors.New("you can add up to 30 selectionProblemAnswer")
	}

	someStatementGroups := lo.GroupBy(selectionProblemAnswers, func(s *SelectionProblemAnswer) string {
		return s.Statement()
	})
	for _, answers := range someStatementGroups {
		if len(answers) > 1 {
			return errors.New("same correct statement is not accepted")
		}
	}

	correctAnswerCount := lo.CountBy(selectionProblemAnswers, func(answer *SelectionProblemAnswer) bool {
		return answer.IsCorrect()
	})
	if correctAnswerCount != 1 {
		return errors.New("only one correct answer")
	}

	s.selectionProblemAnswers = selectionProblemAnswers
	return nil
}

func (s *SelectionProblem) SetStatement(statement selection_problems.Statement) {
	s.statement = statement
}
