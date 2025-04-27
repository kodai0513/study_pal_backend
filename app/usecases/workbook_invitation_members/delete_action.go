package workbook_invitation_members

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	UserId                     uuid.UUID
	WorkbookId                 uuid.UUID
	WorkbookInvitationMemberId uuid.UUID
}

type DeleteAction struct {
	PermissionGuard                    permission_guard.WorkbookPermissionGuard
	Tx                                 trancaction.Tx
	WorkbookInvitationMemberRepository repositories.WorkbookInvitationMemberRepository
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	err := a.PermissionGuard.Check("delete:workbook-invitation-members", command.UserId, command.WorkbookId)
	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}

	existInvitationMember := a.WorkbookInvitationMemberRepository.ExistByIdAndWorkbookId(command.WorkbookInvitationMemberId, command.WorkbookId)
	if !existInvitationMember {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbookInvitationMember not found")))
	}

	err = trancaction.WithTx(a.Tx, func() {
		a.WorkbookInvitationMemberRepository.Delete(command.WorkbookInvitationMemberId, command.WorkbookId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}
	return nil
}
