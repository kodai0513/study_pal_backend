package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/master_datas/master_roles"

	"github.com/google/uuid"
)

type Workbook struct {
	id                 uuid.UUID
	description        workbooks.Description
	isPublic           bool
	problems           []*Problem
	title              workbooks.Title
	userId             uuid.UUID
	workbookCategories []*WorkbookCategory
	workbookMembers    []*WorkbookMember
}

func CreateWorkbook(id uuid.UUID, description workbooks.Description, userId uuid.UUID, title workbooks.Title) *Workbook {
	workbookMembers := []*WorkbookMember{NewWorkbookMember(uuid.New(), master_roles.Admin, userId, id)}

	return &Workbook{
		id:              id,
		isPublic:        false,
		description:     description,
		title:           title,
		userId:          userId,
		workbookMembers: workbookMembers,
	}
}

func NewWorkbook(
	id uuid.UUID,
	isPublic bool,
	description workbooks.Description,
	title workbooks.Title,
	userId uuid.UUID,
	workbookMembers []*WorkbookMember,
) *Workbook {
	return &Workbook{
		id:              id,
		isPublic:        isPublic,
		description:     description,
		title:           title,
		userId:          userId,
		workbookMembers: workbookMembers,
	}
}

func (w *Workbook) AddProblems(problem *Problem) error {
	if problem.workbookCategoryId != uuid.Nil || problem.workbookCategoryClassificationId != uuid.Nil {
		return errors.New("only unclassified questions can be added")
	}

	w.problems = append(w.problems, problem)

	return nil
}

func (w *Workbook) ChangePublic() error {
	if len(w.problems) == 0 && len(w.workbookCategories) == 0 {
		return errors.New("cannot be included without classification")
	}
	w.isPublic = true
	return nil
}

func (w *Workbook) ChangePrivate() {
	w.isPublic = false
}

func (w *Workbook) SetDescription(value workbooks.Description) {
	w.description = value
}

func (w *Workbook) SetTitle(value workbooks.Title) {
	w.title = value
}

func (w *Workbook) Id() uuid.UUID {
	return w.id
}

func (w *Workbook) Description() string {
	return w.description.Value()
}

func (w *Workbook) IsPublic() bool {
	return w.isPublic
}

func (w *Workbook) Title() string {
	return w.title.Value()
}

func (w *Workbook) UserId() uuid.UUID {
	return w.userId
}

func (w *Workbook) WorkbookMembers() []*WorkbookMember {
	return w.workbookMembers
}
