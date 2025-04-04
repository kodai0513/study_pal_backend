package selection_problems

import "github.com/google/uuid"

type SelectionProblemDto struct {
	SelectionProblemAnswers []*SelectionProblemAnswerDto
	Statement               string
}

type SelectionProblemAnswerDto struct {
	IsCorrect                bool
	SelectionProblemAnswerId uuid.UUID
	Statement                string
}
