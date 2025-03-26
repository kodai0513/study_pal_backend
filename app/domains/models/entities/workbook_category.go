package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type WorkbookCategory struct {
	id                      uuid.UUID
	name                    workbook_categories.Name
	problems                map[uuid.UUID]*Problem
	workbookCategoryDetails map[uuid.UUID]*WorkbookCategoryDetail
	workbookId              uuid.UUID
}

func CreateWorkbookCategory(id uuid.UUID, name workbook_categories.Name, workbookId uuid.UUID) *WorkbookCategory {
	return &WorkbookCategory{
		id:         id,
		name:       name,
		workbookId: workbookId,
	}
}

func NewWorkbookCategory(id uuid.UUID, name workbook_categories.Name, problems []*Problem, workbookCategoryDetails []*WorkbookCategoryDetail, workbookId uuid.UUID) *WorkbookCategory {
	return &WorkbookCategory{
		id:                      id,
		name:                    name,
		problems:                lo.SliceToMap(problems, func(p *Problem) (uuid.UUID, *Problem) { return p.Id(), p }),
		workbookCategoryDetails: lo.SliceToMap(workbookCategoryDetails, func(w *WorkbookCategoryDetail) (uuid.UUID, *WorkbookCategoryDetail) { return w.Id(), w }),
		workbookId:              workbookId,
	}
}

func (w *WorkbookCategory) Id() uuid.UUID {
	return w.id
}

func (w *WorkbookCategory) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategory) Problems() map[uuid.UUID]*Problem {
	return w.problems
}

func (w *WorkbookCategory) WorkbookCategoryDetails() map[uuid.UUID]*WorkbookCategoryDetail {
	return w.workbookCategoryDetails
}

func (w *WorkbookCategory) WorkbookId() uuid.UUID {
	return w.workbookId
}

func (w *WorkbookCategory) SetName(name workbook_categories.Name) {
	w.name = name
}

func (w *WorkbookCategory) AddProblem(problem *Problem) error {
	if problem.WorkbookCategoryId() != w.id || problem.WorkbookCategoryDetailId() != uuid.Nil {
		return errors.New("only category questions can be added")
	}

	w.problems[problem.Id()] = problem

	return nil
}

func (w *WorkbookCategory) AddWorkbookCategoryDetail(workbookCategoryDetail *WorkbookCategoryDetail) error {
	if w.workbookId != workbookCategoryDetail.WorkbookCategoryId() {
		return errors.New("cannot include that category in detail")
	}

	if len(workbookCategoryDetail.Problems()) == 0 {
		return errors.New("more than one issue is required to add category details")
	}

	w.workbookCategoryDetails[workbookCategoryDetail.Id()] = workbookCategoryDetail

	return nil
}
