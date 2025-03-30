package entities

import (
	true_or_false_problem "study-pal-backend/app/domains/models/value_objects/true_or_false_problems"

	"github.com/google/uuid"
)

type TrueOrFalseProblem struct {
	id                       uuid.UUID
	isCorrect                bool
	statement                true_or_false_problem.Statement
	workbookCategoryDetailId uuid.UUID
	workbookCategoryId       uuid.UUID
	workbookId               uuid.UUID
}

func NewTrueOrFalseProblem(
	id uuid.UUID,
	isCorrect bool,
	workbookCategoryDetailId uuid.UUID,
	workbookCategoryId uuid.UUID,
	workbookId uuid.UUID,
) *TrueOrFalseProblem {
	return &TrueOrFalseProblem{
		id:                       id,
		isCorrect:                isCorrect,
		workbookCategoryDetailId: workbookCategoryDetailId,
		workbookCategoryId:       workbookCategoryId,
		workbookId:               workbookId,
	}
}

func (t *TrueOrFalseProblem) Id() uuid.UUID {
	return t.id
}

func (t *TrueOrFalseProblem) IsCorrect() bool {
	return t.isCorrect
}

func (t *TrueOrFalseProblem) Statement() string {
	return t.statement.Value()
}

func (t *TrueOrFalseProblem) WorkbookCategoryDetailId() uuid.UUID {
	return t.workbookCategoryDetailId
}

func (t *TrueOrFalseProblem) WorkbookCategoryId() uuid.UUID {
	return t.workbookCategoryId
}

func (t *TrueOrFalseProblem) WorkbookId() uuid.UUID {
	return t.workbookId
}

func (t *TrueOrFalseProblem) SetIsCorrect(isCorrect bool) {
	t.isCorrect = isCorrect
}

func (t *TrueOrFalseProblem) SetStatement(statement true_or_false_problem.Statement) {
	t.statement = statement
}
