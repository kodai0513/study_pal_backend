package answer_types

import (
	"study-pal-backend/app/master_datas/master_answer_types"

	"github.com/google/uuid"
)

type AnswerTypeId struct {
	value uuid.UUID
}

func DescriptionAnswerTypeId() AnswerTypeId {
	return AnswerTypeId{value: master_answer_types.Description}
}

func MultiChoiceAnswerTypeId() AnswerTypeId {
	return AnswerTypeId{value: master_answer_types.MultiChoice}
}

func TruthAnswerTypeId() AnswerTypeId {
	return AnswerTypeId{value: master_answer_types.Truth}
}

func NewAnswerTypeId(value uuid.UUID) AnswerTypeId {
	return AnswerTypeId{value: value}
}

func (a *AnswerTypeId) Value() uuid.UUID {
	return a.value
}
