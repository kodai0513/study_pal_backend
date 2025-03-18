package answer_multi_choices

import "study-pal-backend/app/domains/models/shared"

type AnswerMultiChoiceId struct {
	value int
}

func NewAnswerMultiChoiceId(value int) (AnswerMultiChoiceId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return AnswerMultiChoiceId{value: 0}, err
	}
	return AnswerMultiChoiceId{value: id.Value()}, nil
}

func (a *AnswerMultiChoiceId) Value() int {
	return a.value
}
