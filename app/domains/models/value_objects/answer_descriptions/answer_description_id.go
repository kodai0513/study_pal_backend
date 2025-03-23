package answer_descriptions

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type AnswerDescriptionId struct {
	value uuid.UUID
}

func CreateAnswerDescriptionId() AnswerDescriptionId {
	id := ids.CreateId()
	return AnswerDescriptionId{value: id.Value()}
}

func NewAnswerDescriptionId(value uuid.UUID) AnswerDescriptionId {
	return AnswerDescriptionId{value: value}
}

func (a *AnswerDescriptionId) Value() uuid.UUID {
	return a.value
}
