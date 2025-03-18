package answer_truths

import "study-pal-backend/app/domains/models/shared"

type AnswerTruthId struct {
	value int
}

func NewAnswerTruthId(value int) (AnswerTruthId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return AnswerTruthId{value: 0}, err
	}

	return AnswerTruthId{value: id.Value()}, nil
}

func (a *AnswerTruthId) Value() int {
	return a.value
}
