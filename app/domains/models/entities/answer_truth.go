package entities

import (
	"study-pal-backend/app/domains/models/value_objects/answer_truths"
	"study-pal-backend/app/domains/models/value_objects/problems"

	"github.com/google/uuid"
)

type AnswerTruth struct {
	id        answer_truths.AnswerTruthId
	problemId problems.ProblemId
	truth     answer_truths.Truth
}

func NewAnswerTruth(id answer_truths.AnswerTruthId, problemId problems.ProblemId, truth answer_truths.Truth) *AnswerTruth {
	return &AnswerTruth{
		id:        id,
		problemId: problemId,
		truth:     truth,
	}
}

func (a *AnswerTruth) Id() uuid.UUID {
	return a.id.Value()
}

func (a *AnswerTruth) ProblemId() uuid.UUID {
	return a.problemId.Value()
}

func (a *AnswerTruth) Truth() bool {
	return a.truth.Value()
}
