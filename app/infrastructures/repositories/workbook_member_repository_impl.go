package repositories

import (
	"context"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
)

type WorkbookMemberRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewWorkbookMemberRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.WorkbookMemberRepository {
	return &WorkbookMemberRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}