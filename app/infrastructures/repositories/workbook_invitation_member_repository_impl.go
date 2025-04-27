package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/workbookinvitationmember"

	"github.com/google/uuid"
)

type WorkbookInvitationMemberRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewWorkbookInvitationMemberRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.WorkbookInvitationMemberRepository {
	return &WorkbookInvitationMemberRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (w *WorkbookInvitationMemberRepositoryImpl) Create(invitationMember *entities.WorkbookInvitationMember) {
	w.tx.WorkbookInvitationMember.Create().
		SetID(invitationMember.Id()).
		SetEffectiveAt(invitationMember.EffectiveAt()).
		SetIsInvited(invitationMember.IsInvited()).
		SetRoleID(invitationMember.RoleId()).
		SetUserID(invitationMember.UserId()).
		SetWorkbookID(invitationMember.WorkbookId()).
		SaveX(w.ctx)
}

func (w *WorkbookInvitationMemberRepositoryImpl) Delete(id uuid.UUID, workbookId uuid.UUID) {
	w.tx.WorkbookInvitationMember.DeleteOneID(id).Where(workbookinvitationmember.WorkbookIDNEQ(workbookId)).ExecX(w.ctx)
}

func (w *WorkbookInvitationMemberRepositoryImpl) ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool {
	return w.tx.WorkbookInvitationMember.Query().
		Where(
			workbookinvitationmember.IDEQ(id),
			workbookinvitationmember.WorkbookIDEQ(workbookId),
		).
		ExistX(w.ctx)

}

func (w *WorkbookInvitationMemberRepositoryImpl) FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.WorkbookInvitationMember {
	invitationMember := w.tx.WorkbookInvitationMember.Query().
		Where(
			workbookinvitationmember.IDEQ(id),
			workbookinvitationmember.WorkbookIDEQ(workbookId),
		).
		FirstX(w.ctx)

	if invitationMember == nil {
		return nil
	}

	return entities.NewWorkbookInvitationMember(
		invitationMember.ID,
		invitationMember.EffectiveAt,
		invitationMember.IsInvited,
		invitationMember.RoleID,
		invitationMember.UserID,
		invitationMember.WorkbookID,
	)
}

func (w *WorkbookInvitationMemberRepositoryImpl) Update(invitationMember *entities.WorkbookInvitationMember, workbookId uuid.UUID) {
	w.tx.WorkbookInvitationMember.UpdateOneID(invitationMember.Id()).
		Where(workbookinvitationmember.WorkbookIDEQ(workbookId)).
		SetEffectiveAt(invitationMember.EffectiveAt()).
		SetIsInvited(invitationMember.IsInvited()).
		SetRoleID(invitationMember.RoleId()).
		SetUserID(invitationMember.UserId()).
		SetWorkbookID(invitationMember.WorkbookId()).
		SaveX(w.ctx)
}
