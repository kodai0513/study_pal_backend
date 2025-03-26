package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/master_datas/master_roles"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Workbook struct {
	id                 uuid.UUID
	description        workbooks.Description
	isPublic           bool
	problems           map[uuid.UUID]*Problem
	title              workbooks.Title
	userId             uuid.UUID
	workbookCategories map[uuid.UUID]*WorkbookCategory
	workbookMembers    map[uuid.UUID]*WorkbookMember
}

func CreateWorkbook(id uuid.UUID, description workbooks.Description, userId uuid.UUID, title workbooks.Title) *Workbook {
	newMemberId := uuid.New()
	workbookMembers := map[uuid.UUID]*WorkbookMember{
		newMemberId: NewWorkbookMember(uuid.New(), master_roles.Admin, userId, id),
	}

	return &Workbook{
		id:                 id,
		isPublic:           false,
		description:        description,
		problems:           make(map[uuid.UUID]*Problem, 0),
		title:              title,
		userId:             userId,
		workbookCategories: make(map[uuid.UUID]*WorkbookCategory, 0),
		workbookMembers:    workbookMembers,
	}
}

func NewWorkbook(
	id uuid.UUID,
	description workbooks.Description,
	isPublic bool,
	problems []*Problem,
	title workbooks.Title,
	userId uuid.UUID,
	workbookCategories []*WorkbookCategory,
	workbookMembers []*WorkbookMember,
) *Workbook {
	return &Workbook{
		id:                 id,
		description:        description,
		isPublic:           isPublic,
		problems:           lo.SliceToMap(problems, func(w *Problem) (uuid.UUID, *Problem) { return w.Id(), w }),
		title:              title,
		userId:             userId,
		workbookCategories: lo.SliceToMap(workbookCategories, func(w *WorkbookCategory) (uuid.UUID, *WorkbookCategory) { return w.Id(), w }),
		workbookMembers:    lo.SliceToMap(workbookMembers, func(w *WorkbookMember) (uuid.UUID, *WorkbookMember) { return w.Id(), w }),
	}
}

func (w *Workbook) AddProblems(problem *Problem) error {
	if problem.WorkbookCategoryId() != uuid.Nil || problem.WorkbookCategoryDetailId() != uuid.Nil {
		return errors.New("only unclassified questions can be added")
	}

	w.problems[problem.Id()] = problem

	return nil
}

func (w *Workbook) ChangePublic() error {
	if len(w.problems) == 0 && len(w.workbookCategories) == 0 {
		return errors.New("cannot be included without problem")
	}

	w.isPublic = true

	return nil
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

func (w *Workbook) WorkbookMembers() map[uuid.UUID]*WorkbookMember {
	return w.workbookMembers
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
