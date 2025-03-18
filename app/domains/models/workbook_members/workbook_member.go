package workbook_members

import (
	"study-pal-backend/app/domains/models/roles"
	"study-pal-backend/app/domains/models/users"
)

type WorkbookMember struct {
	workbookMemberId WorkbookMemberId
	roleId           roles.RoleId
	userId           users.UserId
}

func NewWorkbookMember(workbookMemberId WorkbookMemberId, roleId roles.RoleId, userId users.UserId) *WorkbookMember {
	return &WorkbookMember{
		workbookMemberId: workbookMemberId,
		roleId:           roleId,
		userId:           userId,
	}
}

func (w *WorkbookMember) WorkbookMemberId() int {
	return w.workbookMemberId.Value()
}

func (w *WorkbookMember) RoleId() int {
	return w.roleId.Value()
}

func (w *WorkbookMember) UserId() int {
	return w.userId.Value()
}
