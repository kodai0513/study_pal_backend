package answer_types

import "study-pal-backend/app/domains/models/shared"

type AnswerTypeId struct {
	value int
}

func NewAnswerTypeId(value int) (AnswerTypeId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return AnswerTypeId{value: 0}, err
	}
	return AnswerTypeId{value: id.Value()}, nil
}

func (a *AnswerTypeId) Value() int {
	return a.value
}
