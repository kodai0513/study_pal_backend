package problems

import "study-pal-backend/app/domains/models/shared"

type ProblemId struct {
	value int
}

func NewProblemId(value int) (ProblemId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return ProblemId{value: 0}, err
	}
	return ProblemId{value: id.Value()}, nil
}

func (p *ProblemId) Value() int {
	return p.value
}
