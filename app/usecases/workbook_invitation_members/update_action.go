package workbook_invitation_members

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	RoleId                     uuid.UUID
	UserId                     uuid.UUID
	WorkbookId                 uuid.UUID
	WorkbookInvitationMemberId uuid.UUID
}

type UpdateAction struct {
	PermissionGuard                    permission_guard.WorkbookPermissionGuard
	Tx                                 trancaction.Tx
	WorkbookInvitationMemberRepository repositories.WorkbookInvitationMemberRepository
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) usecase_error.UsecaseErrorGroup {
	err := a.PermissionGuard.Check("delete:workbook-invitation-members", command.UserId, command.WorkbookId)
	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}

	invitationMember := a.WorkbookInvitationMemberRepository.FindByIdAndWorkbookId(command.WorkbookInvitationMemberId, command.WorkbookId)
	if invitationMember == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbookInvitationMember not found")))
	}

	invitationMember.SetRoleId(command.RoleId)

	err = trancaction.WithTx(a.Tx, func() {
		a.WorkbookInvitationMemberRepository.Update(invitationMember, command.WorkbookId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
