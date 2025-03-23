package entities

import (
	"study-pal-backend/app/domains/models/value_objects/roles"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/models/value_objects/workbook_members"
	"study-pal-backend/app/domains/models/value_objects/workbooks"

	"github.com/google/uuid"
)

type Workbook struct {
	id              workbooks.WorkbookId
	description     workbooks.Description
	isPublic        workbooks.IsPublic
	title           workbooks.Title
	userId          users.UserId
	workbookMembers []*WorkbookMember
}

func CreateWorkbook(id workbooks.WorkbookId, description workbooks.Description, userId users.UserId, title workbooks.Title) *Workbook {
	workbookMemberId := workbook_members.CreateWorkbookMemberId()
	adminRoleId := roles.AdminRoleId()
	workbookMembers := []*WorkbookMember{NewWorkbookMember(workbookMemberId, adminRoleId, userId, id)}

	return &Workbook{
		id:              id,
		isPublic:        workbooks.NewIsPublic(false),
		description:     description,
		title:           title,
		userId:          userId,
		workbookMembers: workbookMembers,
	}
}

func NewWorkbook(
	id workbooks.WorkbookId,
	isPublic workbooks.IsPublic,
	description workbooks.Description,
	title workbooks.Title,
	userId users.UserId,
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

func (w *Workbook) ChangePublic(publishableWorkbook bool) error {
	w.isPublic = workbooks.NewIsPublic(true)
	return nil
}

func (w *Workbook) Id() uuid.UUID {
	return w.id.Value()
}

func (w *Workbook) Description() string {
	return w.description.Value()
}

func (w *Workbook) IsPublic() bool {
	return w.isPublic.Value()
}

func (w *Workbook) Title() string {
	return w.title.Value()
}

func (w *Workbook) UserId() uuid.UUID {
	return w.userId.Value()
}

func (w *Workbook) WorkbookMembers() []*WorkbookMember {
	return w.workbookMembers
}
