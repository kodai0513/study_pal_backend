package entities

import (
	"time"

	"github.com/google/uuid"
)

type WorkbookInvitationMember struct {
	id          uuid.UUID
	effectiveAt time.Time
	isInvited   bool
	roleId      uuid.UUID
	userId      uuid.UUID
	workbookId  uuid.UUID
}

func NewWorkbookInvitationMember(
	id uuid.UUID,
	effectiveAt time.Time,
	isInvited bool,
	roleId uuid.UUID,
	userId uuid.UUID,
	workbookId uuid.UUID,
) *WorkbookInvitationMember {
	return &WorkbookInvitationMember{
		id:          id,
		effectiveAt: effectiveAt,
		isInvited:   isInvited,
		roleId:      roleId,
		userId:      userId,
		workbookId:  workbookId,
	}
}

func (w *WorkbookInvitationMember) Id() uuid.UUID {
	return w.id
}

func (w *WorkbookInvitationMember) EffectiveAt() time.Time {
	return w.effectiveAt
}

func (w *WorkbookInvitationMember) IsInvited() bool {
	return w.isInvited
}

func (w *WorkbookInvitationMember) RoleId() uuid.UUID {
	return w.roleId
}

func (w *WorkbookInvitationMember) UserId() uuid.UUID {
	return w.userId
}

func (w *WorkbookInvitationMember) WorkbookId() uuid.UUID {
	return w.workbookId
}

func (w *WorkbookInvitationMember) ChangeInvited() {
	w.isInvited = true
}

func (w *WorkbookInvitationMember) SetRoleId(roleId uuid.UUID) {
	w.roleId = roleId
}
