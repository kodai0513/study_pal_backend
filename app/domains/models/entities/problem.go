package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/answer_types"
	"study-pal-backend/app/domains/models/value_objects/problems"
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"
	"study-pal-backend/app/domains/models/value_objects/workbook_category_classifications"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/master_datas/master_answer_types"

	"github.com/google/uuid"
)

type Problem struct {
	id                               problems.ProblemId
	answerDescription                *AnswerDescription
	answerMultiChoices               []*AnswerMultiChoice
	answerTruth                      *AnswerTruth
	answerTypeId                     answer_types.AnswerTypeId
	statement                        problems.Statement
	workbookCategoryClassificationId workbook_category_classifications.WorkbookCategoryClassificationId
	workbookCategoryId               workbook_categories.WorkbookCategoryId
	workbookId                       workbooks.WorkbookId
}

func NewProblem(id problems.ProblemId, answerTypeId answer_types.AnswerTypeId, statement problems.Statement, workbookId workbooks.WorkbookId) *Problem {
	return &Problem{
		id:                               id,
		answerTypeId:                     answerTypeId,
		statement:                        statement,
		workbookCategoryClassificationId: workbook_category_classifications.NewWorkbookCategoryClassificationId(uuid.UUID{}),
		workbookCategoryId:               workbook_categories.NewWorkbookCategoryId(uuid.UUID{}),
		workbookId:                       workbookId,
	}
}

func (p *Problem) AddAnswerMultiChoice(answerMultiChoice *AnswerMultiChoice) error {
	if p.answerTypeId.Value() != master_answer_types.MultiChoice {
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
	if p.answerTypeId.Value() != master_answer_types.Description {
		return errors.New("cannot add that answer type")
	}

	p.answerDescription = answerDescription
	return nil
}

func (p *Problem) SetAnswerTruth(answerTruth *AnswerTruth) error {
	if p.answerTypeId.Value() != master_answer_types.Truth {
		return errors.New("cannot add that answer type")
	}

	p.answerTruth = answerTruth
	return nil
}

func (p *Problem) SetWorkbookCategoryClassificationId(workbookCategoryClassificationId workbook_category_classifications.WorkbookCategoryClassificationId) {
	p.workbookCategoryClassificationId = workbookCategoryClassificationId
}

func (p *Problem) SetWorkbookCategoryId(workbookCategoryId workbook_categories.WorkbookCategoryId) {
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
	return p.answerTypeId.Value() == master_answer_types.Description
}

func (p *Problem) IsAnswerTypeMultiChoice() bool {
	return p.answerTypeId.Value() == master_answer_types.MultiChoice
}

func (p *Problem) IsAnswerTypeTruth() bool {
	return p.answerTypeId.Value() == master_answer_types.Truth
}

func (p *Problem) AnswerTypeId() uuid.UUID {
	return p.answerTypeId.Value()
}

func (p *Problem) Id() uuid.UUID {
	return p.id.Value()
}

func (p *Problem) Statement() string {
	return p.statement.Value()
}

func (p *Problem) WorkbookCategoryClassificationId() uuid.UUID {
	return p.workbookCategoryClassificationId.Value()
}

func (p *Problem) WorkbookCategoryId() uuid.UUID {
	return p.workbookCategoryId.Value()
}

func (p *Problem) WorkbookId() uuid.UUID {
	return p.workbookId.Value()
}
