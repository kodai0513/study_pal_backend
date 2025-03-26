package entities

import (
	"github.com/google/uuid"
)

type AnswerTruth struct {
	id        uuid.UUID
	problemId uuid.UUID
	truth     bool
}

func NewAnswerTruth(id uuid.UUID, problemId uuid.UUID, truth bool) *AnswerTruth {
	return &AnswerTruth{
		id:        id,
		problemId: problemId,
		truth:     truth,
	}
}

func (a *AnswerTruth) Id() uuid.UUID {
	return a.id
}

func (a *AnswerTruth) ProblemId() uuid.UUID {
	return a.problemId
}

func (a *AnswerTruth) Truth() bool {
	return a.truth
}

func (a *AnswerTruth) SetTruth(truth bool) {
	a.truth = truth
}
