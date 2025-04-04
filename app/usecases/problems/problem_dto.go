package problems

import "github.com/google/uuid"

type DescriptionProblemDto struct {
	CorrentStatement         string
	Statement                string
	WorkbookCategoryDetailId *uuid.UUID
	WorkbookCategoryId       *uuid.UUID
}

type SelectionProblemDto struct {
	SelectionProblemAnswers  []*SelectionProblemAnswerDto
	Statement                string
	WorkbookCategoryDetailId *uuid.UUID
	WorkbookCategoryId       *uuid.UUID
}

type SelectionProblemAnswerDto struct {
	IsCorrect bool
	Statement string
}

type TrueOrFalseProblemDto struct {
	IsCorrect                bool
	Statement                string
	WorkbookCategoryDetailId *uuid.UUID
	WorkbookCategoryId       *uuid.UUID
}

type ProblemDto struct {
	DescriptionProblemDtos []*DescriptionProblemDto
	SelectionProblemDtos   []*SelectionProblemDto
	TrueOrFalseProblemDtos []*TrueOrFalseProblemDto
}
