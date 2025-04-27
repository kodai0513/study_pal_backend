package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type WorkbookInvitationMemberRepository interface {
	Create(invitationMember *entities.WorkbookInvitationMember)
	Delete(id uuid.UUID, workbookId uuid.UUID)
	ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool
	FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.WorkbookInvitationMember
	Update(invitationMember *entities.WorkbookInvitationMember, workbookId uuid.UUID)
}
