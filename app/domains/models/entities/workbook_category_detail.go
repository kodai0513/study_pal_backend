package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"

	"github.com/google/uuid"
)

type WorkbookCategoryDetail struct {
	id                 uuid.UUID
	name               workbook_categories.Name
	problems           map[uuid.UUID]*Problem
	workbookCategoryId uuid.UUID
}

func NewWorkbookCategoryDetail(
	id uuid.UUID,
	name workbook_categories.Name,
	workbookCategoryId uuid.UUID,
) *WorkbookCategoryDetail {
	return &WorkbookCategoryDetail{
		id:                 id,
		name:               name,
		problems:           make(map[uuid.UUID]*Problem, 0),
		workbookCategoryId: workbookCategoryId,
	}
}

func (w *WorkbookCategoryDetail) Id() uuid.UUID {
	return w.id
}

func (w *WorkbookCategoryDetail) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategoryDetail) Problems() map[uuid.UUID]*Problem {
	return w.problems
}

func (w *WorkbookCategoryDetail) WorkbookCategoryId() uuid.UUID {
	return w.workbookCategoryId
}

func (w *WorkbookCategoryDetail) SetName(name workbook_categories.Name) {
	w.name = name
}

func (w *WorkbookCategoryDetail) AddProblem(problem *Problem) error {
	if problem.WorkbookCategoryId() != w.workbookCategoryId || problem.WorkbookCategoryDetailId() != w.id {
		return errors.New("that issue cannot be included in the category details")
	}

	w.problems[problem.Id()] = problem

	return nil
}
