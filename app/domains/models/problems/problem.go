package problems

import "study-pal-backend/app/domains/models/workbooks"

type Problem struct {
	problemId  ProblemId
	statement  Statement
	workbookId workbooks.WorkbookId
}

func NewProblem(problemId ProblemId, statement Statement, workbookId workbooks.WorkbookId) *Problem {
	return &Problem{
		problemId:  problemId,
		statement:  statement,
		workbookId: workbookId,
	}
}

func (p *Problem) ProblemId() int {
	return p.problemId.Value()
}

func (p *Problem) Statement() string {
	return p.statement.Value()
}

func (p *Problem) WorkbookId() int {
	return p.workbookId.Value()
}
