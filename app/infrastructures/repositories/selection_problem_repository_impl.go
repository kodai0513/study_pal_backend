package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/selectionproblem"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SelectionProblemRepositoryImpl struct {
	client *ent.Client
	ctx    context.Context
}

func NewSelectionProblemRepositoryImpl(client *ent.Client, ctx context.Context) repositories.SelectionProblemRepository {
	return &SelectionProblemRepositoryImpl{
		client: client,
		ctx:    ctx,
	}
}

func (s *SelectionProblemRepositoryImpl) CreateBulk(problems []*entities.SelectionProblem) []*entities.SelectionProblem {
	problemCreateQueries := make([]*ent.SelectionProblemCreate, 0)
	problemAnswerCreateQueries := make([]*ent.SelectionProblemAnswerCreate, 0)
	for _, p := range problems {
		problemCreateQueries = append(problemCreateQueries,
			s.client.SelectionProblem.Create().
				SetID(p.Id()).
				SetStatement(p.Statement()).
				SetNillableWorkbookCategoryDetailID(p.WorkbookCategoryDetailId()).
				SetNillableWorkbookCategoryID(p.WorkbookCategoryId()).
				SetWorkbookID(p.WorkbookId()),
		)

		for _, pa := range p.SelectionProblemAnswers() {
			problemAnswerCreateQueries = append(problemAnswerCreateQueries,
				s.client.SelectionProblemAnswer.Create().
					SetID(pa.Id()).
					SetIsCorrect(pa.IsCorrect()).
					SetSelectionProblemID(pa.SelectionProblemId()).
					SetStatement(pa.Statement()),
			)
		}
	}

	resultProblems := s.client.SelectionProblem.CreateBulk(problemCreateQueries...).SaveX(s.ctx)
	resultProblemAnswers := s.client.SelectionProblemAnswer.CreateBulk(problemAnswerCreateQueries...).SaveX(s.ctx)
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
			problem.WorkbookCategoryDetailID,
			problem.WorkbookCategoryID,
			problem.WorkbookID,
		)
	})
}

func (s *SelectionProblemRepositoryImpl) Delete(id uuid.UUID) {
	s.client.SelectionProblem.DeleteOneID(id).ExecX(s.ctx)
}

func (s *SelectionProblemRepositoryImpl) FindById(id uuid.UUID) *entities.SelectionProblem {
	result := s.client.SelectionProblem.Query().
		Where(selectionproblem.IDEQ(id)).
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
		result.WorkbookCategoryDetailID,
		result.WorkbookCategoryID,
		result.WorkbookID,
	)
}

func (s *SelectionProblemRepositoryImpl) ExistById(id uuid.UUID) bool {
	return s.client.SelectionProblem.Query().Where(selectionproblem.IDEQ(id)).ExistX(s.ctx)
}

func (s *SelectionProblemRepositoryImpl) Update(problem *entities.SelectionProblem) *entities.SelectionProblem {
	problemResult := s.client.SelectionProblem.UpdateOneID(problem.Id()).
		SetStatement(problem.Statement()).
		SaveX(s.ctx)

	answerResults := s.client.SelectionProblemAnswer.MapCreateBulk(problem.SelectionProblemAnswers(), func(c *ent.SelectionProblemAnswerCreate, i int) {
		c.SetIsCorrect(problem.SelectionProblemAnswers()[i].IsCorrect()).SetStatement(problem.SelectionProblemAnswers()[i].Statement())
	}).SaveX(s.ctx)

	answers := lo.Map(answerResults, func(answer *ent.SelectionProblemAnswer, _ int) *entities.SelectionProblemAnswer {
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
		problemResult.WorkbookCategoryDetailID,
		problemResult.WorkbookCategoryID,
		problemResult.WorkbookID,
	)
}
