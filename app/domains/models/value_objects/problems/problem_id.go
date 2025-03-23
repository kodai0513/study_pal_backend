package problems

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type ProblemId struct {
	value uuid.UUID
}

func CreateProblemId() ProblemId {
	id := ids.CreateId()
	return ProblemId{value: id.Value()}
}

func NewProblemId(value uuid.UUID) ProblemId {
	return ProblemId{value: value}
}

func (p *ProblemId) Value() uuid.UUID {
	return p.value
}
