package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/problems"
	"study-pal-backend/app/master_datas/master_answer_types"

	"github.com/google/uuid"
)

type Problem struct {
	id                               uuid.UUID
	answerDescription                *AnswerDescription
	answerMultiChoices               []*AnswerMultiChoice
	answerTruth                      *AnswerTruth
	answerTypeId                     uuid.UUID
	statement                        problems.Statement
	workbookCategoryClassificationId uuid.UUID
	workbookCategoryId               uuid.UUID
	workbookId                       uuid.UUID
}

func NewProblem(id uuid.UUID, answerTypeId uuid.UUID, statement problems.Statement, workbookId uuid.UUID) *Problem {
	return &Problem{
		id:                               id,
		answerTypeId:                     answerTypeId,
		statement:                        statement,
		workbookCategoryClassificationId: uuid.Nil,
		workbookCategoryId:               uuid.Nil,
		workbookId:                       workbookId,
	}
}

func (p *Problem) AddAnswerMultiChoice(answerMultiChoice *AnswerMultiChoice) error {
	if p.answerTypeId != master_answer_types.MultiChoice {
		return errors.New("cannot add that answer type")
	}
	if len(p.answerMultiChoices) > 30 {
		return errors.New("the upper limit of answerMultiChoice is 30")
	}
	if answerMultiChoice.IsCorrect() {
		for _, answer := range p.answerMultiChoices {
			if answer.IsCorrect() {
				return errors.New("only one correct answer")
			}
		}
	}

	p.answerMultiChoices = append(p.answerMultiChoices, answerMultiChoice)
	return nil
}

func (p *Problem) SetAnswerDescription(answerDescription *AnswerDescription) error {
	if p.answerTypeId != master_answer_types.Description {
		return errors.New("cannot add that answer type")
	}

	p.answerDescription = answerDescription
	return nil
}

func (p *Problem) SetAnswerTruth(answerTruth *AnswerTruth) error {
	if p.answerTypeId != master_answer_types.Truth {
		return errors.New("cannot add that answer type")
	}

	p.answerTruth = answerTruth
	return nil
}

func (p *Problem) SetWorkbookCategoryClassificationId(workbookCategoryClassificationId uuid.UUID) {
	p.workbookCategoryClassificationId = workbookCategoryClassificationId
}

func (p *Problem) SetWorkbookCategoryId(workbookCategoryId uuid.UUID) {
	p.workbookCategoryId = workbookCategoryId
}

func (p *Problem) AnswerDescription() *AnswerDescription {
	return p.answerDescription
}

func (p *Problem) AnswerMultiChoices() []*AnswerMultiChoice {
	return p.answerMultiChoices
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

func (p *Problem) WorkbookCategoryClassificationId() uuid.UUID {
	return p.workbookCategoryClassificationId
}

func (p *Problem) WorkbookCategoryId() uuid.UUID {
	return p.workbookCategoryId
}

func (p *Problem) WorkbookId() uuid.UUID {
	return p.workbookId
}
