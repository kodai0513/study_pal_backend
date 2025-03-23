package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
)

type WorkbookRepositoryImpl struct {
	ctx    context.Context
	client *ent.Client
}

func NewWorkbookRepositoryImpl(ctx context.Context, client *ent.Client) repositories.WorkbookRepository {
	return &WorkbookRepositoryImpl{
		ctx:    ctx,
		client: client,
	}
}

func (p *WorkbookRepositoryImpl) Create(workbook *entities.Workbook) {
	p.client.Workbook.Create().
		SetID(workbook.Id()).
		SetCreatedID(workbook.UserId()).
		SetDescription(workbook.Description()).
		SetIsPublic(workbook.IsPublic()).
		SetTitle(workbook.Title()).
		SaveX(p.ctx)

	p.client.WorkbookMember.Create().
		SetID(workbook.WorkbookMembers()[0].Id()).
		SetRoleID(workbook.WorkbookMembers()[0].RoleId()).
		SetMemberID(workbook.WorkbookMembers()[0].UserId()).
		SetWorkbookID(workbook.WorkbookMembers()[0].WorkbookId()).
		SaveX(p.ctx)
}
