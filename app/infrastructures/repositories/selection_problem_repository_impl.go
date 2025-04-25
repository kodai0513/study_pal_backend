package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/selectionproblem"
	"study-pal-backend/ent/selectionproblemanswer"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SelectionProblemRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewSelectionProblemRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.SelectionProblemRepository {
	return &SelectionProblemRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (s *SelectionProblemRepositoryImpl) CreateBulk(problems []*entities.SelectionProblem) []*entities.SelectionProblem {
	problemCreateQueries := make([]*ent.SelectionProblemCreate, 0)
	problemAnswerCreateQueries := make([]*ent.SelectionProblemAnswerCreate, 0)
	for _, p := range problems {
		problemCreateQueries = append(problemCreateQueries,
			s.tx.SelectionProblem.Create().
				SetID(p.Id()).
				SetStatement(p.Statement()).
				SetNillableWorkbookCategoryID(p.WorkbookCategoryId()).
				SetWorkbookID(p.WorkbookId()),
		)

		for _, pa := range p.SelectionProblemAnswers() {
			problemAnswerCreateQueries = append(problemAnswerCreateQueries,
				s.tx.SelectionProblemAnswer.Create().
					SetID(pa.Id()).
					SetIsCorrect(pa.IsCorrect()).
					SetSelectionProblemID(pa.SelectionProblemId()).
					SetStatement(pa.Statement()),
			)
		}
	}

	resultProblems := s.tx.SelectionProblem.CreateBulk(problemCreateQueries...).SaveX(s.ctx)
	resultProblemAnswers := s.tx.SelectionProblemAnswer.CreateBulk(problemAnswerCreateQueries...).SaveX(s.ctx)
	resultProblemAnswerGroups := make(map[uuid.UUID][]*ent.SelectionProblemAnswer, 0)
	for _, r := range resultProblemAnswers {
		resultProblemAnswerGroups[r.SelectionProblemID] = append(resultProblemAnswerGroups[r.SelectionProblemID], r)
	}
	return lo.Map(resultProblems, func(problem *ent.SelectionProblem, _ int) *entities.SelectionProblem {
		resultAnswers := resultProblemAnswerGroups[problem.ID]
		answers := lo.Map(resultAnswers, func(answer *ent.SelectionProblemAnswer, _ int) *entities.SelectionProblemAnswer {
			statement, _ := selection_problem_answers.NewStatement(answer.Statement)
			return entities.NewSelectionProblemAnswer(
				answer.ID,
				answer.IsCorrect,
				answer.SelectionProblemID,
				statement,
			)
		})
		statement, _ := selection_problems.NewStatement(problem.Statement)
		return entities.NewSelectionProblem(
			problem.ID,
			answers,
			statement,
			problem.WorkbookCategoryID,
			problem.WorkbookID,
		)
	})
}

func (s *SelectionProblemRepositoryImpl) Delete(id uuid.UUID, workbookId uuid.UUID) {
	s.tx.SelectionProblem.DeleteOneID(id).
		Where(selectionproblem.WorkbookIDEQ(workbookId)).
		ExecX(s.ctx)
}

func (s *SelectionProblemRepositoryImpl) FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.SelectionProblem {
	result := s.tx.SelectionProblem.Query().
		Where(
			selectionproblem.IDEQ(id),
			selectionproblem.WorkbookIDEQ(workbookId),
		).
		WithSelectionProblemAnswers().
		FirstX(s.ctx)

	if result == nil {
		return nil
	}

	answers := lo.Map(result.Edges.SelectionProblemAnswers, func(answer *ent.SelectionProblemAnswer, _ int) *entities.SelectionProblemAnswer {
		statement, _ := selection_problem_answers.NewStatement(answer.Statement)
		return entities.NewSelectionProblemAnswer(
			answer.ID,
			answer.IsCorrect,
			answer.SelectionProblemID,
			statement,
		)
	})

	statement, _ := selection_problems.NewStatement(result.Statement)
	return entities.NewSelectionProblem(
		result.ID,
		answers,
		statement,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}

func (s *SelectionProblemRepositoryImpl) ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool {
	return s.tx.SelectionProblem.Query().
		Where(
			selectionproblem.IDEQ(id),
			selectionproblem.WorkbookIDEQ(workbookId),
		).
		ExistX(s.ctx)
}

func (s *SelectionProblemRepositoryImpl) Update(problem *entities.SelectionProblem, workbookId uuid.UUID) *entities.SelectionProblem {
	problemResult := s.tx.SelectionProblem.UpdateOneID(problem.Id()).
		Where(selectionproblem.WorkbookIDEQ(workbookId)).
		SetStatement(problem.Statement()).
		SaveX(s.ctx)

	s.tx.SelectionProblemAnswer.Delete().
		Where(
			selectionproblemanswer.SelectionProblemIDEQ(problem.Id()),
		).
		ExecX(s.ctx)

	createdAnswerResults := s.tx.SelectionProblemAnswer.MapCreateBulk(problem.SelectionProblemAnswers(), func(c *ent.SelectionProblemAnswerCreate, i int) {
		c.
			SetID(problem.SelectionProblemAnswers()[i].Id()).
			SetIsCorrect(problem.SelectionProblemAnswers()[i].IsCorrect()).
			SetSelectionProblemID(problem.SelectionProblemAnswers()[i].SelectionProblemId()).
			SetStatement(problem.SelectionProblemAnswers()[i].Statement())
	}).SaveX(s.ctx)

	answers := lo.Map(createdAnswerResults, func(answer *ent.SelectionProblemAnswer, _ int) *entities.SelectionProblemAnswer {
		statement, _ := selection_problem_answers.NewStatement(answer.Statement)
		return entities.NewSelectionProblemAnswer(
			answer.ID,
			answer.IsCorrect,
			answer.SelectionProblemID,
			statement,
		)
	})

	statement, _ := selection_problems.NewStatement(problemResult.Statement)
	return entities.NewSelectionProblem(
		problemResult.ID,
		answers,
		statement,
		problemResult.WorkbookCategoryID,
		problemResult.WorkbookID,
	)
}
