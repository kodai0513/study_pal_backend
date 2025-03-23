package answer_multi_choices

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type AnswerMultiChoiceId struct {
	value uuid.UUID
}

func CreateAnswerMultiChoiceId() AnswerMultiChoiceId {
	id := ids.CreateId()
	return AnswerMultiChoiceId{value: id.Value()}
}

func NewAnswerMultiChoiceId(value uuid.UUID) AnswerMultiChoiceId {
	return AnswerMultiChoiceId{value: value}
}

func (a *AnswerMultiChoiceId) Value() uuid.UUID {
	return a.value
}
