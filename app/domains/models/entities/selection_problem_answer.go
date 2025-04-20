package entities

import (
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"

	"github.com/google/uuid"
)

type SelectionProblemAnswer struct {
	id                 uuid.UUID
	isCorrect          bool
	selectionProblemId uuid.UUID
	statement          selection_problem_answers.Statement
}

func NewSelectionProblemAnswer(
	id uuid.UUID,
	isCorrect bool,
	selectionProblemId uuid.UUID,
	statement selection_problem_answers.Statement,
) *SelectionProblemAnswer {
	return &SelectionProblemAnswer{
		id:                 id,
		isCorrect:          isCorrect,
		selectionProblemId: selectionProblemId,
		statement:          statement,
	}
}

func (s *SelectionProblemAnswer) Id() uuid.UUID {
	return s.id
}

func (s *SelectionProblemAnswer) IsCorrect() bool {
	return s.isCorrect
}

func (s *SelectionProblemAnswer) SelectionProblemId() uuid.UUID {
	return s.selectionProblemId
}

func (s *SelectionProblemAnswer) Statement() string {
	return s.statement.Value()
}

func (s *SelectionProblemAnswer) SetIsCorrect(isCorrect bool) {
	s.isCorrect = isCorrect
}

func (s *SelectionProblemAnswer) SetStatement(statement selection_problem_answers.Statement) {
	s.statement = statement
}
