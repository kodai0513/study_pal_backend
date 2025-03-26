package entities

import (
	"study-pal-backend/app/domains/models/value_objects/answer_multi_choices"

	"github.com/google/uuid"
)

type AnswerMultiChoice struct {
	id        uuid.UUID
	isCorrect bool
	name      answer_multi_choices.Name
	problemId uuid.UUID
}

func NewAnswerMultiChoice(
	id uuid.UUID,
	isCorrect bool,
	name answer_multi_choices.Name,
	problemId uuid.UUID,
) *AnswerMultiChoice {
	return &AnswerMultiChoice{
		id:        id,
		name:      name,
		isCorrect: isCorrect,
		problemId: problemId,
	}
}

func (a *AnswerMultiChoice) Id() uuid.UUID {
	return a.id
}

func (a *AnswerMultiChoice) IsCorrect() bool {
	return a.isCorrect
}

func (a *AnswerMultiChoice) Name() string {
	return a.name.Value()
}

func (a *AnswerMultiChoice) ProblemId() uuid.UUID {
	return a.problemId
}

func (a *AnswerMultiChoice) SetIsCorrect(isCorrect bool) {
	a.isCorrect = isCorrect
}

func (a *AnswerMultiChoice) SetName(name answer_multi_choices.Name) {
	a.name = name
}
