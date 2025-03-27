package entities

import (
	"study-pal-backend/app/domains/models/value_objects/answer_descriptions"

	"github.com/google/uuid"
)

type AnswerDescription struct {
	id        uuid.UUID
	name      answer_descriptions.Name
	problemId uuid.UUID
}

func NewAnswerDescription(id uuid.UUID, name answer_descriptions.Name, problemId uuid.UUID) *AnswerDescription {
	return &AnswerDescription{
		id:        id,
		name:      name,
		problemId: problemId,
	}
}

func (a *AnswerDescription) Id() uuid.UUID {
	return a.id
}

func (a *AnswerDescription) Name() string {
	return a.name.Value()
}

func (a *AnswerDescription) ProblemId() uuid.UUID {
	return a.problemId
}

func (a *AnswerDescription) setName(name answer_descriptions.Name) {
	a.name = name
}
