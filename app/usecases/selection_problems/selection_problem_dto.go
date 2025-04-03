package selection_problems

type SelectionProblemDto struct {
	SelectionProblemAnswers []*SelectionProblemAnswerDto
	Statement               string
}

type SelectionProblemAnswerDto struct {
	IsCorrect bool
	Statement string
}
