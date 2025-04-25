package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/true_or_false_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/trueorfalseproblem"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type TrueOrFalseProblemRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewTrueOrFalseProblemRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.TrueOrFalseProblemRepository {
	return &TrueOrFalseProblemRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (t *TrueOrFalseProblemRepositoryImpl) CreateBulk(problems []*entities.TrueOrFalseProblem) []*entities.TrueOrFalseProblem {
	results := t.tx.TrueOrFalseProblem.MapCreateBulk(problems, func(tofpc *ent.TrueOrFalseProblemCreate, i int) {
		tofpc.SetID(problems[i].Id()).
			SetIsCorrect(problems[i].IsCorrect()).
			SetStatement(problems[i].Statement()).
			SetNillableWorkbookCategoryID(problems[i].WorkbookCategoryId()).
			SetWorkbookID(problems[i].WorkbookId())
	}).SaveX(t.ctx)

	return lo.Map(results, func(result *ent.TrueOrFalseProblem, _ int) *entities.TrueOrFalseProblem {
		satement, _ := true_or_false_problems.NewStatement(result.Statement)
		return entities.NewTrueOrFalseProblem(
			result.ID,
			result.IsCorrect,
			satement,
			result.WorkbookCategoryID,
			result.WorkbookID,
		)
	})
}

func (t *TrueOrFalseProblemRepositoryImpl) Delete(id uuid.UUID, workbookId uuid.UUID) {
	t.tx.TrueOrFalseProblem.
		DeleteOneID(id).
		Where(trueorfalseproblem.WorkbookIDEQ(workbookId)).
		ExecX(t.ctx)
}

func (t *TrueOrFalseProblemRepositoryImpl) ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool {
	return t.tx.TrueOrFalseProblem.Query().
		Where(
			trueorfalseproblem.IDEQ(id),
			trueorfalseproblem.WorkbookIDEQ(workbookId),
		).
		ExistX(t.ctx)
}

func (t *TrueOrFalseProblemRepositoryImpl) FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.TrueOrFalseProblem {
	result := t.tx.TrueOrFalseProblem.Query().
		Where(
			trueorfalseproblem.IDEQ(id),
			trueorfalseproblem.WorkbookIDEQ(workbookId),
		).
		FirstX(t.ctx)

	if result == nil {
		return nil
	}

	statement, _ := true_or_false_problems.NewStatement(result.Statement)
	return entities.NewTrueOrFalseProblem(
		result.ID,
		result.IsCorrect,
		statement,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}

func (t *TrueOrFalseProblemRepositoryImpl) Update(problem *entities.TrueOrFalseProblem, workbookId uuid.UUID) *entities.TrueOrFalseProblem {
	result := t.tx.TrueOrFalseProblem.UpdateOneID(problem.Id()).
		Where(trueorfalseproblem.WorkbookIDEQ(workbookId)).
		SetIsCorrect(problem.IsCorrect()).
		SetStatement(problem.Statement()).
		SaveX(t.ctx)

	statement, _ := true_or_false_problems.NewStatement(result.Statement)
	return entities.NewTrueOrFalseProblem(
		result.ID,
		result.IsCorrect,
		statement,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}
