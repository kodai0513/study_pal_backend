package entities

import (
	"study-pal-backend/app/domains/models/value_objects/answer_descriptions"
	"study-pal-backend/app/domains/models/value_objects/problems"

	"github.com/google/uuid"
)

type AnswerDescription struct {
	id        answer_descriptions.AnswerDescriptionId
	name      answer_descriptions.Name
	problemId problems.ProblemId
}

func NewAnswerDescription(id answer_descriptions.AnswerDescriptionId, name answer_descriptions.Name, problemId problems.ProblemId) *AnswerDescription {
	return &AnswerDescription{
		id:        id,
		name:      name,
		problemId: problemId,
	}
}

func (a *AnswerDescription) Id() uuid.UUID {
	return a.id.Value()
}

func (a *AnswerDescription) Name() string {
	return a.name.Value()
}

func (a *AnswerDescription) ProblemId() uuid.UUID {
	return a.problemId.Value()
}
