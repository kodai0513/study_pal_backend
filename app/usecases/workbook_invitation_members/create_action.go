package workbook_invitation_members

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"
	"time"

	"github.com/google/uuid"
)

type CreateActionCommand struct {
	RoleId     uuid.UUID
	UserId     uuid.UUID
	WorkbookId uuid.UUID
}

type CreateAction struct {
	PermissionGuard                    permission_guard.WorkbookPermissionGuard
	Tx                                 trancaction.Tx
	WorkbookInvitationMemberRepository repositories.WorkbookInvitationMemberRepository
}

func (a *CreateAction) Execute(command *CreateActionCommand) usecase_error.UsecaseErrorGroup {
	err := a.PermissionGuard.Check("create:workbook-invitation-members", command.UserId, command.WorkbookId)
	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}

	err = trancaction.WithTx(a.Tx, func() {
		a.WorkbookInvitationMemberRepository.Create(entities.NewWorkbookInvitationMember(
			uuid.New(),
			time.Now().AddDate(0, 0, 7),
			false,
			command.RoleId,
			command.UserId,
			command.WorkbookId,
		))
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
