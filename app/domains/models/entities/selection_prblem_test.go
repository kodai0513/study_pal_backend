package entities

import (
	"fmt"
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSelectionProblem_ReplaceSelectionProblemAnswer_正常系とエラー系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryId,
		uuid.New(),
	)

	// 正常系のテスト: 正常な回答セットの作成
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1 := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement1)

	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2 := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerStatement2)

	newAnswers := []*SelectionProblemAnswer{answer1, answer2}

	// 正常に回答を置き換えられることを確認
	err := problem.ReplaceSelectionProblemAnswer(newAnswers)
	assert.NoError(t, err)
	assert.Len(t, problem.SelectionProblemAnswers(), 2)
	assert.Equal(t, newAnswers, problem.SelectionProblemAnswers())

	// エラー系テスト1: 回答数が足りない場合
	insufficientAnswers := []*SelectionProblemAnswer{answer1}
	err = problem.ReplaceSelectionProblemAnswer(insufficientAnswers)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "at least 2 selectionProblemAnswer are required")

	// エラー系テスト2: 同じ文章の回答がある場合
	answerDuplicateStatement, _ := selection_problem_answers.NewStatement("回答1")
	answerDuplicate := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerDuplicateStatement)
	duplicateAnswers := []*SelectionProblemAnswer{answer1, answerDuplicate}

	err = problem.ReplaceSelectionProblemAnswer(duplicateAnswers)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "same correct statement is not accepted")

	// エラー系テスト3: 複数の正解回答がある場合
	answerStatement3, _ := selection_problem_answers.NewStatement("回答3")
	answer3 := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement3)
	multipleCorrectAnswers := []*SelectionProblemAnswer{answer1, answer2, answer3}

	err = problem.ReplaceSelectionProblemAnswer(multipleCorrectAnswers)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only one correct answer")

	// エラー系テスト4: 回答数が多すぎる場合（31個）
	tooManyAnswers := make([]*SelectionProblemAnswer, 0)
	tooManyAnswers = append(tooManyAnswers, answer1)

	// 30個を超える回答を作成
	for i := 0; i < 30; i++ {
		statementText := fmt.Sprintf("回答%d", i+2)
		answerStatement, _ := selection_problem_answers.NewStatement(statementText)
		answer := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerStatement)
		tooManyAnswers = append(tooManyAnswers, answer)
	}

	err = problem.ReplaceSelectionProblemAnswer(tooManyAnswers)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "you can add up to 30 selectionProblemAnswer")

	// 確認: 元の回答セットが影響を受けていないことを確認
	assert.Len(t, problem.SelectionProblemAnswers(), 2)
	assert.Equal(t, newAnswers, problem.SelectionProblemAnswers())
}
