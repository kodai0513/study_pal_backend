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
	client *ent.Client
	ctx    context.Context
}

func NewTrueOrFalseProblemRepositoryImpl(client *ent.Client, ctx context.Context) repositories.TrueOrFalseProblemRepository {
	return &TrueOrFalseProblemRepositoryImpl{
		client: client,
		ctx:    ctx,
	}
}

func (t *TrueOrFalseProblemRepositoryImpl) CreateBulk(problems []*entities.TrueOrFalseProblem) []*entities.TrueOrFalseProblem {
	results := t.client.TrueOrFalseProblem.MapCreateBulk(problems, func(tofpc *ent.TrueOrFalseProblemCreate, i int) {
		tofpc.SetID(problems[i].Id()).
			SetIsCorrect(problems[i].IsCorrect()).
			SetStatement(problems[i].Statement()).
			SetNillableWorkbookCategoryDetailID(problems[i].WorkbookCategoryDetailId()).
			SetNillableWorkbookCategoryID(problems[i].WorkbookCategoryId()).
			SetWorkbookID(problems[i].WorkbookId())
	}).SaveX(t.ctx)

	return lo.Map(results, func(result *ent.TrueOrFalseProblem, _ int) *entities.TrueOrFalseProblem {
		satement, _ := true_or_false_problems.NewStatement(result.Statement)
		return entities.NewTrueOrFalseProblem(
			result.ID,
			result.IsCorrect,
			satement,
			result.WorkbookCategoryDetailID,
			result.WorkbookCategoryID,
			result.WorkbookID,
		)
	})
}

func (t *TrueOrFalseProblemRepositoryImpl) Delete(id uuid.UUID) {
	t.client.TrueOrFalseProblem.DeleteOneID(id).ExecX(t.ctx)
}

func (t *TrueOrFalseProblemRepositoryImpl) ExistById(id uuid.UUID) bool {
	return t.client.TrueOrFalseProblem.Query().Where(trueorfalseproblem.IDEQ(id)).ExistX(t.ctx)
}

func (t *TrueOrFalseProblemRepositoryImpl) FindById(id uuid.UUID) *entities.TrueOrFalseProblem {
	result := t.client.TrueOrFalseProblem.Query().Where(trueorfalseproblem.IDEQ(id)).FirstX(t.ctx)

	if result == nil {
		return nil
	}

	statement, _ := true_or_false_problems.NewStatement(result.Statement)
	return entities.NewTrueOrFalseProblem(
		result.ID,
		result.IsCorrect,
		statement,
		result.WorkbookCategoryDetailID,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}

func (t *TrueOrFalseProblemRepositoryImpl) Update(problem *entities.TrueOrFalseProblem) *entities.TrueOrFalseProblem {
	result := t.client.TrueOrFalseProblem.UpdateOneID(problem.Id()).
		SetIsCorrect(problem.IsCorrect()).
		SetStatement(problem.Statement()).
		SaveX(t.ctx)

	statement, _ := true_or_false_problems.NewStatement(result.Statement)
	return entities.NewTrueOrFalseProblem(
		result.ID,
		result.IsCorrect,
		statement,
		result.WorkbookCategoryDetailID,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}
