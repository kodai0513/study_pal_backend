package answer_truths

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type AnswerTruthId struct {
	value uuid.UUID
}

func CreateAnswerTruthId() AnswerTruthId {
	id := ids.CreateId()
	return AnswerTruthId{value: id.Value()}
}

func NewAnswerTruthId(value uuid.UUID) AnswerTruthId {
	return AnswerTruthId{value: value}
}

func (a *AnswerTruthId) Value() uuid.UUID {
	return a.value
}
