package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/description_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/descriptionproblem"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DescriptionProblemRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewDescriptionProblemRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.DescriptionProblemRepository {
	return &DescriptionProblemRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (d *DescriptionProblemRepositoryImpl) CreateBulk(problems []*entities.DescriptionProblem) []*entities.DescriptionProblem {
	results := d.tx.DescriptionProblem.MapCreateBulk(problems, func(dpc *ent.DescriptionProblemCreate, i int) {
		dpc.SetID(problems[i].Id()).
			SetCorrectStatement(problems[i].CorrectStatement()).
			SetStatement(problems[i].Statement()).
			SetNillableWorkbookCategoryID(problems[i].WorkbookCategoryId()).
			SetWorkbookID(problems[i].WorkbookId())
	}).SaveX(d.ctx)

	return lo.Map(results, func(result *ent.DescriptionProblem, _ int) *entities.DescriptionProblem {
		correctStatement, _ := description_problems.NewCorrectStatement(result.CorrectStatement)
		statement, _ := description_problems.NewStatement(result.Statement)
		return entities.NewDescriptionProblem(
			result.ID,
			correctStatement,
			statement,
			result.WorkbookCategoryID,
			result.WorkbookID,
		)
	})
}

func (d *DescriptionProblemRepositoryImpl) Delete(id uuid.UUID, workbookId uuid.UUID) {
	d.tx.DescriptionProblem.DeleteOneID(id).
		Where(descriptionproblem.WorkbookIDEQ(workbookId)).
		ExecX(d.ctx)
}

func (d *DescriptionProblemRepositoryImpl) ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool {
	return d.tx.DescriptionProblem.Query().
		Where(
			descriptionproblem.IDEQ(id),
			descriptionproblem.WorkbookIDEQ(workbookId),
		).
		ExistX(d.ctx)
}

func (d *DescriptionProblemRepositoryImpl) ExistByWorkbookId(workbookId uuid.UUID) bool {
	return d.tx.DescriptionProblem.Query().Where(descriptionproblem.IDEQ(workbookId)).ExistX(d.ctx)
}

func (d *DescriptionProblemRepositoryImpl) FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.DescriptionProblem {
	result := d.tx.DescriptionProblem.Query().
		Where(
			descriptionproblem.IDEQ(id),
			descriptionproblem.WorkbookIDEQ(workbookId),
		).
		FirstX(d.ctx)

	if result == nil {
		return nil
	}

	correctStatement, _ := description_problems.NewCorrectStatement(result.CorrectStatement)
	statement, _ := description_problems.NewStatement(result.Statement)
	return entities.NewDescriptionProblem(
		result.ID,
		correctStatement,
		statement,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}

func (d *DescriptionProblemRepositoryImpl) Update(problem *entities.DescriptionProblem, workbookId uuid.UUID) *entities.DescriptionProblem {
	result := d.tx.DescriptionProblem.
		UpdateOneID(problem.Id()).
		Where(descriptionproblem.WorkbookIDEQ(workbookId)).
		SetCorrectStatement(problem.CorrectStatement()).
		SetStatement(problem.Statement()).
		SaveX(d.ctx)

	correctStatement, _ := description_problems.NewCorrectStatement(result.CorrectStatement)
	statement, _ := description_problems.NewStatement(result.Statement)
	return entities.NewDescriptionProblem(
		result.ID,
		correctStatement,
		statement,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}
