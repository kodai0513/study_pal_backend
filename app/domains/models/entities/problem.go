package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/problems"
	"study-pal-backend/app/master_datas/master_answer_types"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Problem struct {
	id                       uuid.UUID
	answerDescription        *AnswerDescription
	answerMultiChoices       map[uuid.UUID]*AnswerMultiChoice
	answerTruth              *AnswerTruth
	answerTypeId             uuid.UUID
	statement                problems.Statement
	workbookCategoryDetailId uuid.UUID
	workbookCategoryId       uuid.UUID
	workbookId               uuid.UUID
}

func CreateProblem(
	id uuid.UUID,
	answerDescription *AnswerDescription,
	answerMultiChoices []*AnswerMultiChoice,
	answerTruth *AnswerTruth,
	answerTypeId uuid.UUID,
	statement problems.Statement,
	workbookCategoryDetailId uuid.UUID,
	workbookCategoryId uuid.UUID,
	workbookId uuid.UUID,
) *Problem {
	return &Problem{
		id:                       id,
		answerDescription:        answerDescription,
		answerMultiChoices:       lo.SliceToMap(answerMultiChoices, func(a *AnswerMultiChoice) (uuid.UUID, *AnswerMultiChoice) { return a.Id(), a }),
		answerTypeId:             answerTypeId,
		statement:                statement,
		workbookCategoryDetailId: workbookCategoryDetailId,
		workbookCategoryId:       workbookCategoryId,
		workbookId:               workbookId,
	}
}

func NewProblem(id uuid.UUID, answerTypeId uuid.UUID, statement problems.Statement, workbookId uuid.UUID) *Problem {
	return &Problem{
		id:                       id,
		answerMultiChoices:       make(map[uuid.UUID]*AnswerMultiChoice, 0),
		answerTypeId:             answerTypeId,
		statement:                statement,
		workbookCategoryDetailId: uuid.Nil,
		workbookCategoryId:       uuid.Nil,
		workbookId:               workbookId,
	}
}

func (p *Problem) AnswerDescription() *AnswerDescription {
	return p.answerDescription
}

func (p *Problem) AnswerMultiChoices() []*AnswerMultiChoice {
	return lo.MapToSlice(p.answerMultiChoices, func(k uuid.UUID, a *AnswerMultiChoice) *AnswerMultiChoice {
		return a
	})
}

func (p *Problem) AnswerTruth() *AnswerTruth {
	return p.answerTruth
}

func (p *Problem) IsAnswerTypeDescription() bool {
	return p.answerTypeId == master_answer_types.Description
}

func (p *Problem) IsAnswerTypeMultiChoice() bool {
	return p.answerTypeId == master_answer_types.MultiChoice
}

func (p *Problem) IsAnswerTypeTruth() bool {
	return p.answerTypeId == master_answer_types.Truth
}

func (p *Problem) AnswerTypeId() uuid.UUID {
	return p.answerTypeId
}

func (p *Problem) Id() uuid.UUID {
	return p.id
}

func (p *Problem) Statement() string {
	return p.statement.Value()
}

func (p *Problem) WorkbookCategoryDetailId() uuid.UUID {
	return p.workbookCategoryDetailId
}

func (p *Problem) WorkbookCategoryId() uuid.UUID {
	return p.workbookCategoryId
}

func (p *Problem) WorkbookId() uuid.UUID {
	return p.workbookId
}

func (p *Problem) setAnswerDescription(answerDescription *AnswerDescription) error {
	if p.answerTypeId != master_answer_types.Description {
		return errors.New("cannot add that answer type")
	}

	p.answerDescription = answerDescription
	return nil
}

func (p *Problem) replaceAnswerMultiChoices(answerMultiChoices []*AnswerMultiChoice) error {
	if p.answerTypeId != master_answer_types.MultiChoice {
		return errors.New("cannot add that answer type")
	}

	if len(answerMultiChoices) <= 2 {
		return errors.New("the lower limit is at least 2")
	}

	if len(answerMultiChoices) > 30 {
		return errors.New("the upper limit of answerMultiChoice is 30")
	}

	correctCount := lo.CountBy(
		answerMultiChoices,
		func(answerMultiChoice *AnswerMultiChoice) bool {
			return answerMultiChoice.isCorrect
		},
	)

	if correctCount != 1 {
		return errors.New("only one correct answer can be set")
	}

	p.answerMultiChoices = lo.SliceToMap(
		answerMultiChoices,
		func(answerMultiChoice *AnswerMultiChoice) (uuid.UUID, *AnswerMultiChoice) {
			return answerMultiChoice.Id(), answerMultiChoice
		},
	)
	return nil
}

func (p *Problem) setAnswerTruth(answerTruth *AnswerTruth) error {
	if p.answerTypeId != master_answer_types.Truth {
		return errors.New("cannot add that answer type")
	}

	p.answerTruth = answerTruth
	return nil
}

func (p *Problem) setStatement(statement problems.Statement) {
	p.statement = statement
}

func (p *Problem) setWorkbookCategoryDetailId(workbookCategoryDetailId uuid.UUID) {
	p.workbookCategoryDetailId = workbookCategoryDetailId
}

func (p *Problem) setWorkbookCategoryId(workbookCategoryId uuid.UUID) {
	p.workbookCategoryId = workbookCategoryId
}
