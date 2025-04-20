package entities

import (
	"study-pal-backend/app/domains/models/value_objects/description_problems"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDescriptionProblem_正常系(t *testing.T) {
	// UUIDの生成
	problemId := uuid.New()
	workbookId := uuid.New()
	workbookCategoryId := uuid.New()

	// 問題文と正解の生成
	statement, err := description_problems.NewStatement("これは記述問題です")
	assert.NoError(t, err)

	correctStatement, err := description_problems.NewCorrectStatement("これが正解です")
	assert.NoError(t, err)

	// DescriptionProblemの作成
	problem := NewDescriptionProblem(
		problemId,
		correctStatement,
		statement,
		&workbookCategoryId,
		workbookId,
	)

	// 基本的な属性の確認
	assert.Equal(t, problemId, problem.Id())
	assert.Equal(t, "これは記述問題です", problem.Statement())
	assert.Equal(t, "これが正解です", problem.CorrectStatement())
	assert.Equal(t, workbookId, problem.WorkbookId())
	assert.Equal(t, &workbookCategoryId, problem.WorkbookCategoryId())
}

func TestDescriptionProblem_Statement_正常系(t *testing.T) {
	// 問題文の作成 - 下限値（1文字）
	minStatement, err := description_problems.NewStatement("a")
	assert.NoError(t, err)
	assert.Equal(t, "a", minStatement.Value())

	// 問題文の作成 - 通常値
	normalStatement, err := description_problems.NewStatement("これは通常の長さの問題文です。")
	assert.NoError(t, err)
	assert.Equal(t, "これは通常の長さの問題文です。", normalStatement.Value())

	// 問題文の作成 - 上限値（1000文字）
	longText := generateString(1000)
	maxStatement, err := description_problems.NewStatement(longText)
	assert.NoError(t, err)
	assert.Equal(t, longText, maxStatement.Value())
}

func TestDescriptionProblem_Statement_エラー系(t *testing.T) {
	// 問題文の作成 - 空文字
	_, err := description_problems.NewStatement("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be blank")

	// 問題文の作成 - 上限超過（1001文字）
	longText := generateString(1001)
	_, err = description_problems.NewStatement(longText)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "the length must be between 1 and 1000")
}

func TestDescriptionProblem_CorrectStatement_正常系(t *testing.T) {
	// 正解文の作成 - 下限値（1文字）
	minCorrectStatement, err := description_problems.NewCorrectStatement("a")
	assert.NoError(t, err)
	assert.Equal(t, "a", minCorrectStatement.Value())

	// 正解文の作成 - 通常値
	normalCorrectStatement, err := description_problems.NewCorrectStatement("これは正解の文章です。")
	assert.NoError(t, err)
	assert.Equal(t, "これは正解の文章です。", normalCorrectStatement.Value())

	// 正解文の作成 - 上限値（100文字）
	longText := generateString(100)
	maxCorrectStatement, err := description_problems.NewCorrectStatement(longText)
	assert.NoError(t, err)
	assert.Equal(t, longText, maxCorrectStatement.Value())
}

func TestDescriptionProblem_CorrectStatement_エラー系(t *testing.T) {
	// 正解文の作成 - 空文字
	_, err := description_problems.NewCorrectStatement("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be blank")

	// 正解文の作成 - 上限超過（101文字）
	longText := generateString(101)
	_, err = description_problems.NewCorrectStatement(longText)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "the length must be between 1 and 100")
}

func TestDescriptionProblem_全属性取得確認(t *testing.T) {
	// テスト用データ作成
	problemId := uuid.New()
	workbookId := uuid.New()
	workbookCategoryId := uuid.New()

	statement, _ := description_problems.NewStatement("問題文")
	correctStatement, _ := description_problems.NewCorrectStatement("正解")

	// インスタンス作成
	problem := NewDescriptionProblem(
		problemId,
		correctStatement,
		statement,
		&workbookCategoryId,
		workbookId,
	)

	// 各ゲッターメソッドの動作確認
	assert.Equal(t, problemId, problem.Id())
	assert.Equal(t, "問題文", problem.Statement())
	assert.Equal(t, "正解", problem.CorrectStatement())
	assert.Equal(t, workbookId, problem.WorkbookId())
	assert.Equal(t, &workbookCategoryId, problem.WorkbookCategoryId())
}

// // テスト用ヘルパー関数：指定した長さの文字列を生成
func generateString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[i%len(chars)]
	}
	return string(result)
}
