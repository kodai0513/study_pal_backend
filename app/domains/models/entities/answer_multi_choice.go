package entities

import (
	"study-pal-backend/app/domains/models/value_objects/answer_multi_choices"
	"study-pal-backend/app/domains/models/value_objects/problems"

	"github.com/google/uuid"
)

type AnswerMultiChoice struct {
	id        answer_multi_choices.AnswerMultiChoiceId
	isCorrect answer_multi_choices.IsCorrect
	name      answer_multi_choices.Name
	problemId problems.ProblemId
}

func NewAnswerMultiChoice(
	id answer_multi_choices.AnswerMultiChoiceId,
	isCorrect answer_multi_choices.IsCorrect,
	name answer_multi_choices.Name,
	problemId problems.ProblemId,
) *AnswerMultiChoice {
	return &AnswerMultiChoice{
		id:        id,
		name:      name,
		isCorrect: isCorrect,
		problemId: problemId,
	}
}

func (a *AnswerMultiChoice) Id() uuid.UUID {
	return a.id.Value()
}

func (a *AnswerMultiChoice) IsCorrect() bool {
	return a.isCorrect.Value()
}

func (a *AnswerMultiChoice) Name() string {
	return a.name.Value()
}

func (a *AnswerMultiChoice) ProblemId() uuid.UUID {
	return a.problemId.Value()
}
