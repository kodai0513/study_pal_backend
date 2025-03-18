package answer_multi_choices

import (
	"study-pal-backend/app/domains/models/problems"
)

type AnswerMultiChoice struct {
	answerMultiChoiceId AnswerMultiChoiceId
	isCorrect           IsCorrect
	name                Name
	problemId           problems.ProblemId
}

func NewAnswerMultiChoice(answerMultiChoiceId AnswerMultiChoiceId, isCorrect IsCorrect, name Name, problemId problems.ProblemId) *AnswerMultiChoice {
	return &AnswerMultiChoice{
		answerMultiChoiceId: answerMultiChoiceId,
		name:                name,
		isCorrect:           isCorrect,
		problemId:           problemId,
	}
}

func (a *AnswerMultiChoice) AnswerMultiChoiceId() int {
	return a.answerMultiChoiceId.Value()
}

func (a *AnswerMultiChoice) IsCorrect() bool {
	return a.isCorrect.Value()
}

func (a *AnswerMultiChoice) Name() string {
	return a.name.Value()
}

func (a *AnswerMultiChoice) ProblemId() int {
	return a.problemId.Value()
}
