package answer_descriptions

import "study-pal-backend/app/domains/models/shared"

type AnswerDescriptionId struct {
	value int
}

func NewAnswerDescriptionId(value int) (AnswerDescriptionId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return AnswerDescriptionId{value: 0}, err
	}

	return AnswerDescriptionId{value: id.Value()}, nil
}

func (a *AnswerDescriptionId) Value() int {
	return a.value
}
