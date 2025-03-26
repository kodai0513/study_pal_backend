package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookmember"

	"github.com/google/uuid"
	lo "github.com/samber/lo"
)

type WorkbookRepositoryImpl struct {
	client *ent.Client
	ctx    context.Context
}

func NewWorkbookRepositoryImpl(client *ent.Client, ctx context.Context) repositories.WorkbookRepository {
	return &WorkbookRepositoryImpl{
		client: client,
		ctx:    ctx,
	}
}

func (w *WorkbookRepositoryImpl) Create(workbook *entities.Workbook) *entities.Workbook {
	// workbook 登録
	resultWorkbook := w.client.Workbook.Create().
		SetID(workbook.Id()).
		SetCreatedID(workbook.UserId()).
		SetDescription(workbook.Description()).
		SetIsPublic(workbook.IsPublic()).
		SetTitle(workbook.Title()).
		SaveX(w.ctx)

	// workbookMember 登録
	resultMember := w.client.WorkbookMember.MapCreateBulk(
		workbook.WorkbookMembers(),
		func(wmc *ent.WorkbookMemberCreate)
	).
		SetID(workbook.WorkbookMembers()[0].Id()).
		SetRoleID(workbook.WorkbookMembers()[0].RoleId()).
		SetMemberID(workbook.WorkbookMembers()[0].UserId()).
		SetWorkbookID(workbook.WorkbookMembers()[0].WorkbookId()).
		SaveX(w.ctx)

	member := entities.NewWorkbookMember(resultMember.ID, resultMember.RoleID, resultMember.MemberID, resultMember.WorkbookID)
	description, _ := workbooks.NewDescription(*resultWorkbook.Description)
	title, _ := workbooks.NewTitle(resultWorkbook.Title)
	return entities.NewWorkbook(
		resultWorkbook.ID,
		resultWorkbook.IsPublic,
		description,
		title,
		resultWorkbook.CreatedID,
		[]*entities.WorkbookMember{member},
	)
}

func (w *WorkbookRepositoryImpl) Delete(workbookId uuid.UUID) {
	w.client.Workbook.DeleteOneID(workbookId).ExecX(w.ctx)
}

func (w *WorkbookRepositoryImpl) FindById(workbookId uuid.UUID) *entities.Workbook {
	result := w.client.Workbook.
		Query().
		Where(workbook.IDEQ(workbookId)).
		WithWorkbookMembers().
		FirstX(w.ctx)

	if result == nil {
		return nil
	}

	workbookMembers := lo.Map(
		result.Edges.WorkbookMembers,
		func(member *ent.WorkbookMember, index int) *entities.WorkbookMember {
			return entities.NewWorkbookMember(
				member.ID,
				member.RoleID,
				member.MemberID,
				member.WorkbookID,
			)
		},
	)
	description, _ := workbooks.NewDescription(*result.Description)
	title, _ := workbooks.NewTitle(result.Title)
	return entities.NewWorkbook(
		result.ID,
		result.IsPublic,
		description,
		title,
		result.CreatedID,
		workbookMembers,
	)
}

func (w *WorkbookRepositoryImpl) Update(workbook *entities.Workbook) *entities.Workbook {
	// workbook 更新
	resultWorkbook := w.client.Workbook.UpdateOneID(workbook.Id()).
		SetDescription(workbook.Description()).
		SetTitle(workbook.Title()).
		SetIsPublic(workbook.IsPublic()).
		SaveX(w.ctx)

	// workbookMember 更新
	registerdMemberIds := w.client.WorkbookMember.Query().Where(workbookmember.WorkbookIDEQ(workbook.Id())).IDsX(w.ctx)
	newRegisterMemberIdKeys := lo.SliceToMap(
		workbook.WorkbookMembers(),
		func(workbookMember *entities.WorkbookMember) (uuid.UUID, uuid.UUID) {
			return workbookMember.Id(), workbookMember.Id()
		},
	)
	deleteMemberIds := lo.Filter(registerdMemberIds, func(registerdMemberId uuid.UUID, index int) bool {
		_, ok := newRegisterMemberIdKeys[registerdMemberId]
		return !ok
	})
	w.client.WorkbookMember.Delete().
		Where(
			workbookmember.MemberIDIn(deleteMemberIds...),
			workbookmember.WorkbookIDEQ(workbook.Id()),
		).
		ExecX(w.ctx)

	registerMemberIdKeys := lo.SliceToMap(
		registerdMemberIds,
		func(registerMemberId uuid.UUID) (uuid.UUID, uuid.UUID) {
			return registerMemberId, registerMemberId
		},
	)

	addRegisterMembers := lo.Filter(
		workbook.WorkbookMembers(),
		func(workbookMember *entities.WorkbookMember, index int) bool {
			_, ok := registerMemberIdKeys[workbookMember.Id()]
			return !ok
		},
	)

	resultMembers := w.client.WorkbookMember.MapCreateBulk(
		addRegisterMembers,
		func(wmc *ent.WorkbookMemberCreate, i int) {
			wmc.SetID(addRegisterMembers[i].Id()).
				SetRoleID(addRegisterMembers[i].RoleId()).
				SetMemberID(addRegisterMembers[i].UserId()).
				SetWorkbookID(addRegisterMembers[i].WorkbookId())
		},
	).SaveX(w.ctx)

	members := lo.Map(
		resultMembers,
		func(member *ent.WorkbookMember, index int) *entities.WorkbookMember {
			return entities.NewWorkbookMember(member.ID, member.RoleID, member.MemberID, member.WorkbookID)
		},
	)
	description, _ := workbooks.NewDescription(*resultWorkbook.Description)
	title, _ := workbooks.NewTitle(resultWorkbook.Title)
	return entities.NewWorkbook(
		resultWorkbook.ID,
		resultWorkbook.IsPublic,
		description,
		title,
		resultWorkbook.CreatedID,
		members,
	)
}
