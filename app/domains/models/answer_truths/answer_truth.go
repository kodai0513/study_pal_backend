package answer_truths

import "study-pal-backend/app/domains/models/problems"

type AnswerTruth struct {
	answerTruthId AnswerTruthId
	problemId     problems.ProblemId
	truth         Truth
}

func NewAnswerTruth(answerTruthId AnswerTruthId, problemId problems.ProblemId, truth Truth) *AnswerTruth {
	return &AnswerTruth{
		answerTruthId: answerTruthId,
		problemId:     problemId,
		truth:         truth,
	}
}
