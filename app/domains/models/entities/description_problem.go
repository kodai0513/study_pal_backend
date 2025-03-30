package entities

import (
	"study-pal-backend/app/domains/models/value_objects/description_problems"

	"github.com/google/uuid"
)

type DescriptionProblem struct {
	id                       uuid.UUID
	correntStatement         description_problems.CorrectStatement
	statement                description_problems.Statement
	workbookCategoryDetailId uuid.UUID
	workbookCategoryId       uuid.UUID
	workbookId               uuid.UUID
}

func NewDescriptionProblem(
	id uuid.UUID,
	correctStatement description_problems.CorrectStatement,
	statement description_problems.Statement,
	workbookId uuid.UUID,
	workbookCategoryId uuid.UUID,
	workbookCategoryDetailId uuid.UUID,
) *DescriptionProblem {
	return &DescriptionProblem{
		id:                       id,
		correntStatement:         correctStatement,
		statement:                statement,
		workbookCategoryDetailId: workbookCategoryDetailId,
		workbookCategoryId:       workbookCategoryId,
		workbookId:               workbookId,
	}
}

func (d *DescriptionProblem) Id() uuid.UUID {
	return d.id
}

func (d *DescriptionProblem) CorrectStatement() string {
	return d.correntStatement.Value()
}

func (d *DescriptionProblem) Statement() string {
	return d.statement.Value()
}

func (d *DescriptionProblem) WorkbookCategoryDetailId() uuid.UUID {
	return d.workbookCategoryDetailId
}

func (d *DescriptionProblem) WorkbookCategoryId() uuid.UUID {
	return d.workbookCategoryId
}

func (d *DescriptionProblem) WorkbookId() uuid.UUID {
	return d.workbookId
}
