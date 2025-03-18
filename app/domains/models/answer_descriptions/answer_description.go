package answer_descriptions

import "study-pal-backend/app/domains/models/problems"

type AnswerDescription struct {
	answerDescriptionId AnswerDescriptionId
	name                Name
	problemId           problems.ProblemId
}

func NewAnswerDescription(answerDescriptionId AnswerDescriptionId, name Name, problemId problems.ProblemId) *AnswerDescription {
	return &AnswerDescription{
		answerDescriptionId: answerDescriptionId,
		name:                name,
		problemId:           problemId,
	}
}

func (a *AnswerDescription) AnswerDescriptionId() int {
	return a.answerDescriptionId.Value()
}

func (a *AnswerDescription) Name() string {
	return a.name.Value()
}

func (a *AnswerDescription) ProblemId() int {
	return a.problemId.Value()
}
