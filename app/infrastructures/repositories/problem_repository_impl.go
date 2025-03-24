package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/ent"
	"study-pal-backend/ent/problem"

	"github.com/google/uuid"
)

type ProblemRepositoryImpl struct {
	Client *ent.Client
	Ctx    context.Context
}

func (p *ProblemRepositoryImpl) CreateBulk(problems []*entities.Problem) {
	p.Client.Problem.MapCreateBulk(problems, func(c *ent.ProblemCreate, i int) {
		c.SetID(problems[i].Id()).
			SetAnswerTypeID(problems[i].AnswerTypeId()).
			SetStatement(problems[i].Statement()).
			SetWorkbookID(problems[i].WorkbookId()).
			SetWorkbookCategoryID(problems[i].WorkbookCategoryId()).
			SetWorkbookCategoryClassificationID(problems[i].WorkbookCategoryClassificationId())

		if problems[i].IsAnswerTypeDescription() {
			c.SetAnswerDescriptions(&ent.AnswerDescription{
				ID:        problems[i].AnswerDescription().Id(),
				Name:      problems[i].AnswerDescription().Name(),
				ProblemID: problems[i].AnswerDescription().ProblemId(),
			})
		}
		if problems[i].IsAnswerTypeMultiChoice() {
			var answerMultiChoices []*ent.AnswerMultiChoices
			for _, multiChoice := range problems[i].AnswerMultiChoices() {
				answerMultiChoices = append(
					answerMultiChoices,
					&ent.AnswerMultiChoices{
						ID:        multiChoice.Id(),
						Name:      multiChoice.Name(),
						IsCorrect: multiChoice.IsCorrect(),
						ProblemID: multiChoice.ProblemId(),
					},
				)
			}
			c.AddAnswerMultiChoices(answerMultiChoices...)
		}
		if problems[i].IsAnswerTypeTruth() {
			c.SetAnswerTruths(&ent.AnswerTruth{
				ID:        problems[i].AnswerTruth().Id(),
				ProblemID: problems[i].AnswerDescription().ProblemId(),
				Truth:     problems[i].AnswerTruth().Truth(),
			})
		}
	}).SaveX(p.Ctx)
}

func (p *ProblemRepositoryImpl) ExistByWorkbookId(workbookId uuid.UUID) bool {
	return p.Client.Problem.Query().Where(problem.WorkbookIDEQ(workbookId)).ExistX(p.Ctx)
}
