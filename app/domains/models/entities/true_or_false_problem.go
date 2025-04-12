package entities

import (
	"study-pal-backend/app/domains/models/value_objects/true_or_false_problems"

	"github.com/google/uuid"
)

type TrueOrFalseProblem struct {
	id                 uuid.UUID
	isCorrect          bool
	statement          true_or_false_problems.Statement
	workbookCategoryId *uuid.UUID
	workbookId         uuid.UUID
}

func NewTrueOrFalseProblem(
	id uuid.UUID,
	isCorrect bool,
	statement true_or_false_problems.Statement,
	workbookCategoryId *uuid.UUID,
	workbookId uuid.UUID,
) *TrueOrFalseProblem {
	return &TrueOrFalseProblem{
		id:                 id,
		isCorrect:          isCorrect,
		statement:          statement,
		workbookCategoryId: workbookCategoryId,
		workbookId:         workbookId,
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

func (t *TrueOrFalseProblem) WorkbookCategoryId() *uuid.UUID {
	return t.workbookCategoryId
}

func (t *TrueOrFalseProblem) WorkbookId() uuid.UUID {
	return t.workbookId
}

func (t *TrueOrFalseProblem) SetIsCorrect(isCorrect bool) {
	t.isCorrect = isCorrect
}

func (t *TrueOrFalseProblem) SetStatement(statement true_or_false_problems.Statement) {
	t.statement = statement
}
