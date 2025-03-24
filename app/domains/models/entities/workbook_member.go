package entities

import (
	"study-pal-backend/app/master_datas/master_roles"

	"github.com/google/uuid"
)

type WorkbookMember struct {
	id         uuid.UUID
	roleId     uuid.UUID
	userId     uuid.UUID
	workbookId uuid.UUID
}

func NewWorkbookMember(id uuid.UUID, roleId uuid.UUID, userId uuid.UUID, workbookId uuid.UUID) *WorkbookMember {
	return &WorkbookMember{
		id:         id,
		roleId:     roleId,
		userId:     userId,
		workbookId: workbookId,
	}
}

func (w *WorkbookMember) Id() uuid.UUID {
	return w.id
}

func (w *WorkbookMember) IsAdmin() bool {
	return w.roleId == master_roles.Admin
}

func (w *WorkbookMember) IsEditor() bool {
	return w.roleId == master_roles.Editor
}

func (w *WorkbookMember) IsReader() bool {
	return w.roleId == master_roles.Reader
}

func (w *WorkbookMember) RoleId() uuid.UUID {
	return w.roleId
}

func (w *WorkbookMember) UserId() uuid.UUID {
	return w.userId
}

func (w *WorkbookMember) WorkbookId() uuid.UUID {
	return w.workbookId
}
