package entities

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SelectionProblem struct {
	id                       uuid.UUID
	selectionProblemAnswers  map[uuid.UUID]*SelectionProblemAnswer
	statement                selection_problems.Statement
	workbookCategoryDetailId *uuid.UUID
	workbookCategoryId       *uuid.UUID
	workbookId               uuid.UUID
}

func CreateSelectionProblem(
	id uuid.UUID,
	statement selection_problems.Statement,
	workbookCategoryDetailId *uuid.UUID,
	workbookCategoryId *uuid.UUID,
	workbookId uuid.UUID,
) *SelectionProblem {
	return &SelectionProblem{
		id:                       id,
		selectionProblemAnswers:  make(map[uuid.UUID]*SelectionProblemAnswer, 0),
		statement:                statement,
		workbookCategoryDetailId: workbookCategoryDetailId,
		workbookCategoryId:       workbookCategoryId,
		workbookId:               workbookId,
	}
}

func NewSelectionProblem(
	id uuid.UUID,
	selectionProblemAnswers []*SelectionProblemAnswer,
	statement selection_problems.Statement,
	workbookCategoryDetailId *uuid.UUID,
	workbookCategoryId *uuid.UUID,
	workbookId uuid.UUID,
) *SelectionProblem {
	return &SelectionProblem{
		id: id,
		selectionProblemAnswers: lo.SliceToMap(selectionProblemAnswers, func(selectionProblemAnswer *SelectionProblemAnswer) (uuid.UUID, *SelectionProblemAnswer) {
			return selectionProblemAnswer.Id(), selectionProblemAnswer
		}),
		statement:                statement,
		workbookCategoryDetailId: workbookCategoryDetailId,
		workbookCategoryId:       workbookCategoryId,
		workbookId:               workbookId,
	}
}

func (s *SelectionProblem) Id() uuid.UUID {
	return s.id
}

func (s *SelectionProblem) SelectionProblemAnswers() []*SelectionProblemAnswer {
	return lo.MapToSlice(s.selectionProblemAnswers, func(_ uuid.UUID, s *SelectionProblemAnswer) *SelectionProblemAnswer {
		return s
	})
}

func (s *SelectionProblem) Statement() string {
	return s.statement.Value()
}

func (s *SelectionProblem) WorkbookCategoryDetailId() *uuid.UUID {
	return s.workbookCategoryDetailId
}

func (s *SelectionProblem) WorkbookCategoryId() *uuid.UUID {
	return s.workbookCategoryId
}

func (s *SelectionProblem) WorkbookId() uuid.UUID {
	return s.workbookId
}

func (s *SelectionProblem) AddSelectionProblemAnswer(selectionProblemAnswer *SelectionProblemAnswer) error {
	if len(s.selectionProblemAnswers) >= 30 {
		return errors.New("you can add up to 30 selectionProblemAnswer")
	}

	_, exist := s.selectionProblemAnswers[selectionProblemAnswer.id]
	if exist {
		return errors.New("selectionProblemAnswer already exists")
	}

	searchAnswer := lo.MapToSlice(s.selectionProblemAnswers, func(_ uuid.UUID, s *SelectionProblemAnswer) *SelectionProblemAnswer {
		return s
	})

	answerStatementExist := lo.CountBy(searchAnswer, func(s *SelectionProblemAnswer) bool {
		return s.Statement() == selectionProblemAnswer.Statement()
	}) > 0

	if answerStatementExist {
		return errors.New("selectionProblemAnswerStatement already exists")
	}

	if selectionProblemAnswer.isCorrect {
		answerIsCorrectExists := lo.CountBy(searchAnswer, func(s *SelectionProblemAnswer) bool {
			return s.IsCorrect() == selectionProblemAnswer.IsCorrect()
		}) > 0

		if answerIsCorrectExists {
			return errors.New("selectionProblemAnswerIsCorrect already exists")
		}
	}

	s.selectionProblemAnswers[selectionProblemAnswer.id] = selectionProblemAnswer
	return nil
}

func (s *SelectionProblem) RemoveSelectionProblemAnswer(selectionProblemAnswerId uuid.UUID) error {
	_, exist := s.selectionProblemAnswers[selectionProblemAnswerId]
	if !exist {
		return errors.New("selectionProblemAnswer not found")
	}

	delete(s.selectionProblemAnswers, selectionProblemAnswerId)

	return nil
}

func (s *SelectionProblem) HasMinimumAnswers() error {
	if len(s.selectionProblemAnswers) < 2 {
		return errors.New("at least 2 selectionProblemAnswer are required")
	}

	return nil
}

func (s *SelectionProblem) IsCorrectAnswer() error {
	searchAnswer := lo.MapToSlice(s.selectionProblemAnswers, func(_ uuid.UUID, s *SelectionProblemAnswer) *SelectionProblemAnswer {
		return s
	})

	isNoAnswer := lo.CountBy(searchAnswer, func(s *SelectionProblemAnswer) bool {
		return s.IsCorrect()
	}) == 0

	if isNoAnswer {
		return errors.New("is there a correct answer")
	}

	return nil
}

func (s *SelectionProblem) SetAnswerIsCorrect(answerIsCorrect bool, selectionProblemAnswerId uuid.UUID) error {
	_, exist := s.selectionProblemAnswers[selectionProblemAnswerId]
	if !exist {
		return errors.New("selectionProblemAnswer not exists")
	}

	searchAnswer := lo.MapToSlice(s.selectionProblemAnswers, func(_ uuid.UUID, s *SelectionProblemAnswer) *SelectionProblemAnswer {
		return s
	})

	if answerIsCorrect {
		answerIsCorrectExists := lo.CountBy(searchAnswer, func(s *SelectionProblemAnswer) bool {
			return s.IsCorrect()
		}) > 0

		if answerIsCorrectExists {
			return errors.New("only one answer")
		}
	}

	s.selectionProblemAnswers[selectionProblemAnswerId].setIsCorrect(answerIsCorrect)
	return nil
}

func (s *SelectionProblem) SetAnswerStatement(answerStatement selection_problem_answers.Statement, selectionProblemAnswerId uuid.UUID) error {
	_, exist := s.selectionProblemAnswers[selectionProblemAnswerId]
	if !exist {
		return errors.New("selectionProblemAnswer not exists")
	}

	searchAnswer := lo.MapToSlice(s.selectionProblemAnswers, func(_ uuid.UUID, s *SelectionProblemAnswer) *SelectionProblemAnswer {
		return s
	})

	answerStatementExist := lo.CountBy(searchAnswer, func(s *SelectionProblemAnswer) bool {
		return s.Statement() == answerStatement.Value()
	}) > 0

	if answerStatementExist {
		return errors.New("answerStatement cannot be duplicated")
	}

	s.selectionProblemAnswers[selectionProblemAnswerId].setStatement(answerStatement)
	return nil
}

func (s *SelectionProblem) SetStatement(statement selection_problems.Statement) {
	s.statement = statement
}
