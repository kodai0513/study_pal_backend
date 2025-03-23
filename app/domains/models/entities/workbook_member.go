package entities

import (
	"study-pal-backend/app/domains/models/value_objects/roles"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/models/value_objects/workbook_members"
	"study-pal-backend/app/domains/models/value_objects/workbooks"

	"github.com/google/uuid"
)

type WorkbookMember struct {
	id         workbook_members.WorkbookMemberId
	roleId     roles.RoleId
	userId     users.UserId
	workbookId workbooks.WorkbookId
}

func NewWorkbookMember(id workbook_members.WorkbookMemberId, roleId roles.RoleId, userId users.UserId, workbookId workbooks.WorkbookId) *WorkbookMember {
	return &WorkbookMember{
		id:         id,
		roleId:     roleId,
		userId:     userId,
		workbookId: workbookId,
	}
}

func (w *WorkbookMember) Id() uuid.UUID {
	return w.id.Value()
}

func (w *WorkbookMember) RoleId() uuid.UUID {
	return w.roleId.Value()
}

func (w *WorkbookMember) UserId() uuid.UUID {
	return w.userId.Value()
}

func (w *WorkbookMember) WorkbookId() uuid.UUID {
	return w.workbookId.Value()
}
