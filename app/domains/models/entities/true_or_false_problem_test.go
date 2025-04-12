package entities

import (
	"study-pal-backend/app/domains/models/value_objects/true_or_false_problems"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTrueOrFalseProblem_正常系(t *testing.T) {
	// UUIDの生成
	problemId := uuid.New()
	workbookCategoryId := uuid.New()
	workbookId := uuid.New()

	statement, err := true_or_false_problems.NewStatement("これは〇×問題です")
	assert.NoError(t, err)
	// TrueOrFalseProblemの作成
	problem := NewTrueOrFalseProblem(
		problemId,
		true, // isCorrect = true
		statement,
		&workbookCategoryId,
		workbookId,
	)

	// 基本的な属性の確認
	assert.Equal(t, problemId, problem.Id())
	assert.Equal(t, true, problem.IsCorrect())
	assert.Equal(t, "これは〇×問題です", problem.Statement())
	assert.Equal(t, &workbookCategoryId, problem.WorkbookCategoryId())
	assert.Equal(t, workbookId, problem.WorkbookId())
}

func TestTrueOrFalseProblem_SetIsCorrect(t *testing.T) {
	// TrueOrFalseProblemの作成
	statement, err := true_or_false_problems.NewStatement("これは〇×問題です")
	assert.NoError(t, err)
	problemId := uuid.New()
	workbookCategoryId := uuid.New()
	problem := NewTrueOrFalseProblem(
		problemId,
		true, // 初期値はtrue
		statement,
		&workbookCategoryId,
		uuid.New(),
	)

	// 初期値の確認
	assert.Equal(t, true, problem.IsCorrect())

	// 値を反転
	problem.SetIsCorrect(false)
	assert.Equal(t, false, problem.IsCorrect())

	// 元に戻す
	problem.SetIsCorrect(true)
	assert.Equal(t, true, problem.IsCorrect())
}

func TestTrueOrFalseProblem_Statement_正常系(t *testing.T) {
	// TrueOrFalseProblemの作成
	statement, err := true_or_false_problems.NewStatement("これは通常の長さの〇×問題文です。")
	assert.NoError(t, err)
	workbookCategoryId := uuid.New()
	problem := NewTrueOrFalseProblem(
		uuid.New(),
		true,
		statement,
		&workbookCategoryId,
		uuid.New(),
	)

	// 問題文の作成 - 下限値（1文字）
	minStatement, err := true_or_false_problems.NewStatement("a")
	assert.NoError(t, err)
	problem.SetStatement(minStatement)
	assert.Equal(t, "a", problem.Statement())

	// 問題文の作成 - 通常値
	normalStatement, err := true_or_false_problems.NewStatement("これは通常の長さの〇×問題文です。")
	assert.NoError(t, err)
	problem.SetStatement(normalStatement)
	assert.Equal(t, "これは通常の長さの〇×問題文です。", problem.Statement())

	// 問題文の作成 - 上限値（1000文字）
	longText := generateString(1000)
	maxStatement, err := true_or_false_problems.NewStatement(longText)
	assert.NoError(t, err)
	problem.SetStatement(maxStatement)
	assert.Equal(t, longText, problem.Statement())
}

func TestTrueOrFalseProblem_Statement_エラー系(t *testing.T) {
	// 問題文の作成 - 空文字
	_, err := true_or_false_problems.NewStatement("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be blank")

	// 問題文の作成 - 上限超過（1001文字）
	longText := generateString(1001)
	_, err = true_or_false_problems.NewStatement(longText)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "the length must be between 1 and 1000")
}

func TestTrueOrFalseProblem_全属性取得確認(t *testing.T) {
	// テスト用データ作成
	statement, err := true_or_false_problems.NewStatement("〇×問題の文章")
	assert.NoError(t, err)
	problemId := uuid.New()
	workbookCategoryId := uuid.New()
	workbookId := uuid.New()

	// インスタンス作成
	problem := NewTrueOrFalseProblem(
		problemId,
		false,
		statement,
		&workbookCategoryId,
		workbookId,
	)

	// 各ゲッターメソッドの動作確認
	assert.Equal(t, problemId, problem.Id())
	assert.Equal(t, false, problem.IsCorrect())
	assert.Equal(t, "〇×問題の文章", problem.Statement())
	assert.Equal(t, &workbookCategoryId, problem.WorkbookCategoryId())
	assert.Equal(t, workbookId, problem.WorkbookId())
}
