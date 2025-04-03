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
		SetUserID(workbook.UserId()).
		SetDescription(workbook.Description()).
		SetIsPublic(workbook.IsPublic()).
		SetTitle(workbook.Title()).
		SaveX(w.ctx)

	// workbookMember 登録
	resultMember := w.client.WorkbookMember.Create().
		SetID(workbook.WorkbookMembers()[0].Id()).
		SetRoleID(workbook.WorkbookMembers()[0].RoleId()).
		SetUserID(workbook.WorkbookMembers()[0].UserId()).
		SetWorkbookID(workbook.WorkbookMembers()[0].WorkbookId()).
		SaveX(w.ctx)

	member := entities.NewWorkbookMember(resultMember.ID, resultMember.RoleID, resultMember.UserID, resultMember.WorkbookID)
	description, _ := workbooks.NewDescription(*resultWorkbook.Description)
	title, _ := workbooks.NewTitle(resultWorkbook.Title)
	return entities.NewWorkbook(
		resultWorkbook.ID,
		description,
		make([]uuid.UUID, 0),
		resultWorkbook.IsPublic,
		make([]uuid.UUID, 0),
		title,
		make([]uuid.UUID, 0),
		resultWorkbook.UserID,
		[]*entities.WorkbookMember{member},
	)
}

func (w *WorkbookRepositoryImpl) Delete(workbookId uuid.UUID) {
	w.client.Workbook.DeleteOneID(workbookId).ExecX(w.ctx)
}

func (w *WorkbookRepositoryImpl) ExistById(workbookId uuid.UUID) bool {
	return w.client.Workbook.Query().Where(workbook.IDEQ(workbookId)).ExistX(w.ctx)
}

func (w *WorkbookRepositoryImpl) FindById(workbookId uuid.UUID) *entities.Workbook {
	result := w.client.Workbook.
		Query().
		Where(workbook.IDEQ(workbookId)).
		WithDescriptionProblems().
		WithSelectionProblems().
		WithTrueOrFalseProblems().
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
				member.UserID,
				member.WorkbookID,
			)
		},
	)
	descriptionProblemIds := lo.Map(
		result.Edges.DescriptionProblems,
		func(d *ent.DescriptionProblem, _ int) uuid.UUID {
			return d.ID
		},
	)
	selectionProblemIds := lo.Map(
		result.Edges.SelectionProblems,
		func(d *ent.SelectionProblem, _ int) uuid.UUID {
			return d.ID
		},
	)
	trueOrFalseProblemIds := lo.Map(
		result.Edges.TrueOrFalseProblems,
		func(d *ent.TrueOrFalseProblem, _ int) uuid.UUID {
			return d.ID
		},
	)
	description, _ := workbooks.NewDescription(*result.Description)
	title, _ := workbooks.NewTitle(result.Title)
	return entities.NewWorkbook(
		result.ID,
		description,
		descriptionProblemIds,
		result.IsPublic,
		selectionProblemIds,
		title,
		trueOrFalseProblemIds,
		result.UserID,
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
			workbookmember.UserIDIn(deleteMemberIds...),
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
				SetUserID(addRegisterMembers[i].UserId()).
				SetWorkbookID(addRegisterMembers[i].WorkbookId())
		},
	).SaveX(w.ctx)

	members := lo.Map(
		resultMembers,
		func(member *ent.WorkbookMember, index int) *entities.WorkbookMember {
			return entities.NewWorkbookMember(member.ID, member.RoleID, member.UserID, member.WorkbookID)
		},
	)
	description, _ := workbooks.NewDescription(*resultWorkbook.Description)
	title, _ := workbooks.NewTitle(resultWorkbook.Title)
	return entities.NewWorkbook(
		resultWorkbook.ID,
		description,
		workbook.DescriptionProblemIds(),
		resultWorkbook.IsPublic,
		workbook.SelectionProblemsIds(),
		title,
		workbook.TrueOfFalseProblemIds(),
		resultWorkbook.UserID,
		members,
	)
}
