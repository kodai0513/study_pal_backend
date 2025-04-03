package entities

import (
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSelectionProblem_正常に作成と回答追加ができるか(t *testing.T) {
	// 問題文の作成
	statement, err := selection_problems.NewStatement("これはテスト問題です")
	assert.NoError(t, err)

	// UUIDの生成
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	workbookId := uuid.New()

	// SelectionProblemの作成
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		workbookId,
	)

	// 基本的な属性の確認
	assert.Equal(t, problemId, problem.Id())
	assert.Equal(t, "これはテスト問題です", problem.Statement())
	assert.Equal(t, &workbookCategoryDetailId, problem.WorkbookCategoryDetailId())
	assert.Equal(t, &workbookCategoryId, problem.WorkbookCategoryId())
	assert.Equal(t, workbookId, problem.WorkbookId())
	assert.Empty(t, problem.SelectionProblemAnswers())

	// 回答の作成
	answerStatement1, err := selection_problem_answers.NewStatement("回答1")
	assert.NoError(t, err)

	answer1 := NewSelectionProblemAnswer(
		uuid.New(),
		true,
		problemId,
		answerStatement1,
	)

	// 回答の追加
	err = problem.AddSelectionProblemAnswer(answer1)
	assert.NoError(t, err)
	assert.Len(t, problem.SelectionProblemAnswers(), 1)
}

func TestSelectionProblem_AddSelectionProblemAnswer_正常系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// 複数の回答を追加
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1 := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement1)
	err := problem.AddSelectionProblemAnswer(answer1)
	assert.NoError(t, err)

	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2 := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerStatement2)
	err = problem.AddSelectionProblemAnswer(answer2)
	assert.NoError(t, err)

	// 回答が正しく追加されたことを確認
	answers := problem.SelectionProblemAnswers()
	assert.Len(t, answers, 2)

	// 必須回答数の確認
	err = problem.HasMinimumAnswers()
	assert.NoError(t, err)
}

func TestSelectionProblem_AddSelectionProblemAnswer_エラー系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// 正解の回答を追加
	answerId := uuid.New()
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1 := NewSelectionProblemAnswer(answerId, true, problemId, answerStatement1)
	err := problem.AddSelectionProblemAnswer(answer1)
	assert.NoError(t, err)

	// 同じIDの回答を追加しようとするとエラー
	answer1Duplicate := NewSelectionProblemAnswer(answerId, false, problemId, answerStatement1)
	err = problem.AddSelectionProblemAnswer(answer1Duplicate)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "selectionProblemAnswer already exists")

	// 同じ文章の回答を追加しようとするとエラー
	answerSameStatement, _ := selection_problem_answers.NewStatement("回答1")
	answerDuplicateStatement := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerSameStatement)
	err = problem.AddSelectionProblemAnswer(answerDuplicateStatement)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "selectionProblemAnswerStatement already exists")

	// 別の正解の回答を追加しようとするとエラー
	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	anotherCorrectAnswer := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement2)
	err = problem.AddSelectionProblemAnswer(anotherCorrectAnswer)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "selectionProblemAnswerIsCorrect already exists")

	// 上限（30個）の確認はループで生成するとテストが長くなるためスキップ
}

func TestSelectionProblem_IsOneAnswer_正常系とエラー系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// 正解と不正解の回答を追加
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1 := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement1)
	problem.AddSelectionProblemAnswer(answer1)

	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2 := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerStatement2)
	problem.AddSelectionProblemAnswer(answer2)

	// 正常系: 正解と不正解がある場合はエラーなし
	err := problem.IsCorrectAnswer()
	assert.NoError(t, err)

	// エラー系のテスト準備: すべての回答を正解に設定
	// 既存の回答のisCorrectを変更
	for _, answer := range problem.SelectionProblemAnswers() {
		problem.SetAnswerIsCorrect(false, answer.Id())
	}

	// エラー系: すべての回答が不正解の場合はエラー
	err = problem.IsCorrectAnswer()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is there a correct answer")
}

func TestSelectionProblem_SetAnswerStatement_正常系とエラー系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// 回答を追加
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1Id := uuid.New()
	answer1 := NewSelectionProblemAnswer(answer1Id, true, problemId, answerStatement1)
	problem.AddSelectionProblemAnswer(answer1)

	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2Id := uuid.New()
	answer2 := NewSelectionProblemAnswer(answer2Id, false, problemId, answerStatement2)
	problem.AddSelectionProblemAnswer(answer2)

	// 正常系: 回答文を更新
	newStatement, _ := selection_problem_answers.NewStatement("新しい回答1")
	err := problem.SetAnswerStatement(newStatement, answer1Id)
	assert.NoError(t, err)

	// 回答文が更新されたことを確認
	for _, answer := range problem.SelectionProblemAnswers() {
		if answer.Id() == answer1Id {
			assert.Equal(t, "新しい回答1", answer.Statement())
		}
	}

	// エラー系: 存在しない回答IDで更新しようとするとエラー
	nonExistentId := uuid.New()
	err = problem.SetAnswerStatement(newStatement, nonExistentId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "selectionProblemAnswer not exists")

	// エラー系: 既に存在する回答文で更新しようとするとエラー
	duplicateStatement, _ := selection_problem_answers.NewStatement("回答2")
	err = problem.SetAnswerStatement(duplicateStatement, answer1Id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "answerStatement cannot be duplicated")
}

func TestSelectionProblem_SetAnswerIsCorrect_正常系とエラー系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// 回答を追加
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1Id := uuid.New()
	answer1 := NewSelectionProblemAnswer(answer1Id, true, problemId, answerStatement1)
	problem.AddSelectionProblemAnswer(answer1)

	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2Id := uuid.New()
	answer2 := NewSelectionProblemAnswer(answer2Id, false, problemId, answerStatement2)
	problem.AddSelectionProblemAnswer(answer2)

	// 正常系: 不正解から正解に変更しようとするとエラー（既に正解がある）
	err := problem.SetAnswerIsCorrect(true, answer2Id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only one answer")

	// 正常系: 正解から不正解に変更
	err = problem.SetAnswerIsCorrect(false, answer1Id)
	assert.NoError(t, err)

	// 両方不正解になったことを確認
	for _, answer := range problem.SelectionProblemAnswers() {
		assert.False(t, answer.IsCorrect())
	}

	// これで両方不正解なので、今度は正解に変更できるはず
	err = problem.SetAnswerIsCorrect(true, answer2Id)
	assert.NoError(t, err)

	// エラー系: 存在しない回答IDで更新しようとするとエラー
	nonExistentId := uuid.New()
	err = problem.SetAnswerIsCorrect(true, nonExistentId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "selectionProblemAnswer not exists")
}

func TestSelectionProblem_HasMinimumAnswers_正常系とエラー系(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// エラー系: 回答がない場合
	err := problem.HasMinimumAnswers()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "at least 2 selectionProblemAnswer are required")

	// 1つ目の回答を追加
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1 := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement1)
	problem.AddSelectionProblemAnswer(answer1)

	// エラー系: 回答が1つだけの場合
	err = problem.HasMinimumAnswers()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "at least 2 selectionProblemAnswer are required")

	// 2つ目の回答を追加
	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2 := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerStatement2)
	problem.AddSelectionProblemAnswer(answer2)

	// 正常系: 回答が2つ以上ある場合
	err = problem.HasMinimumAnswers()
	assert.NoError(t, err)
}

func TestSelectionProblem_SetStatement(t *testing.T) {
	// 問題の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := CreateSelectionProblem(
		problemId,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		uuid.New(),
	)

	// 問題文を更新
	newStatement, _ := selection_problems.NewStatement("これは新しいテスト問題です")
	problem.SetStatement(newStatement)

	// 問題文が更新されたことを確認
	assert.Equal(t, "これは新しいテスト問題です", problem.Statement())
}

func TestNewSelectionProblem(t *testing.T) {
	// 問題文の作成
	statement, _ := selection_problems.NewStatement("これはテスト問題です")
	problemId := uuid.New()
	workbookCategoryDetailId := uuid.New()
	workbookCategoryId := uuid.New()
	workbookId := uuid.New()

	// 回答の作成
	answerStatement1, _ := selection_problem_answers.NewStatement("回答1")
	answer1 := NewSelectionProblemAnswer(uuid.New(), true, problemId, answerStatement1)

	answerStatement2, _ := selection_problem_answers.NewStatement("回答2")
	answer2 := NewSelectionProblemAnswer(uuid.New(), false, problemId, answerStatement2)

	// NewSelectionProblemで初期化
	answers := []*SelectionProblemAnswer{answer1, answer2}
	problem := NewSelectionProblem(
		problemId,
		answers,
		statement,
		&workbookCategoryDetailId,
		&workbookCategoryId,
		workbookId,
	)

	// 基本的な属性の確認
	assert.Equal(t, problemId, problem.Id())
	assert.Equal(t, "これはテスト問題です", problem.Statement())
	assert.Equal(t, &workbookCategoryDetailId, problem.WorkbookCategoryDetailId())
	assert.Equal(t, &workbookCategoryId, problem.WorkbookCategoryId())
	assert.Equal(t, workbookId, problem.WorkbookId())
	assert.Len(t, problem.SelectionProblemAnswers(), 2)
}
