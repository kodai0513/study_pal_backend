package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/master_datas/master_roles"

	"github.com/google/uuid"
)

type Workbook struct {
	id                    uuid.UUID
	description           workbooks.Description
	descriptionProblemIds []uuid.UUID
	isPublic              bool
	selectionProblemsIds  []uuid.UUID
	title                 workbooks.Title
	trueOfFalseProblemIds []uuid.UUID
	userId                uuid.UUID
	workbookMembers       []*WorkbookMember
}

func CreateWorkbook(id uuid.UUID, description workbooks.Description, userId uuid.UUID, title workbooks.Title) *Workbook {
	workbookMembers := []*WorkbookMember{NewWorkbookMember(uuid.New(), master_roles.Admin, userId, id)}

	return &Workbook{
		id:                    id,
		description:           description,
		descriptionProblemIds: make([]uuid.UUID, 0),
		isPublic:              false,
		selectionProblemsIds:  make([]uuid.UUID, 0),
		title:                 title,
		trueOfFalseProblemIds: make([]uuid.UUID, 0),
		userId:                userId,
		workbookMembers:       workbookMembers,
	}
}

func NewWorkbook(
	id uuid.UUID,
	description workbooks.Description,
	descriptionProblemIds []uuid.UUID,
	isPublic bool,
	selectionProblemsIds []uuid.UUID,
	title workbooks.Title,
	trueOfFalseProblemIds []uuid.UUID,
	userId uuid.UUID,
	workbookMembers []*WorkbookMember,
) *Workbook {
	return &Workbook{
		id:                    id,
		description:           description,
		descriptionProblemIds: descriptionProblemIds,
		isPublic:              isPublic,
		selectionProblemsIds:  selectionProblemsIds,
		title:                 title,
		trueOfFalseProblemIds: trueOfFalseProblemIds,
		userId:                userId,
		workbookMembers:       workbookMembers,
	}
}

func (w *Workbook) Id() uuid.UUID {
	return w.id
}

func (w *Workbook) Description() string {
	return w.description.Value()
}

func (w *Workbook) DescriptionProblemIds() []uuid.UUID {
	return w.descriptionProblemIds
}

func (w *Workbook) IsPublic() bool {
	return w.isPublic
}

func (w *Workbook) SelectionProblemsIds() []uuid.UUID {
	return w.selectionProblemsIds
}

func (w *Workbook) Title() string {
	return w.title.Value()
}

func (w *Workbook) TrueOfFalseProblemIds() []uuid.UUID {
	return w.trueOfFalseProblemIds
}

func (w *Workbook) UserId() uuid.UUID {
	return w.userId
}

func (w *Workbook) WorkbookMembers() []*WorkbookMember {
	return w.workbookMembers
}

func (w *Workbook) ChangePrivate() {
	w.isPublic = false
}

func (w *Workbook) ChangePublic() error {
	if (len(w.descriptionProblemIds) + len(w.selectionProblemsIds) + len(w.trueOfFalseProblemIds)) == 0 {
		return errors.New("cannot be included without problem")
	}

	w.isPublic = true

	return nil
}

func (w *Workbook) SetDescription(value workbooks.Description) {
	w.description = value
}

func (w *Workbook) SetTitle(value workbooks.Title) {
	w.title = value
}
