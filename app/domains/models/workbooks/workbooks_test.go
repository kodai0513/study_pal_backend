package workbooks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkbookId_正常にWorkbookIdを設定できるか(t *testing.T) {
	id, err := NewWorkbookId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestWorkbookId_WorkbookIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewWorkbookId(0)

	assert.Error(t, err)
}

func TestDescription_正常にDescriptionが設定できるか(t *testing.T) {
	input := "テスト説明"
	description, err := NewDescription(input)

	assert.NoError(t, err)
	assert.Equal(t, input, description.Value())
}

func TestDescription_Descriptionが未入力のときエラーになるか(t *testing.T) {
	input := ""
	_, err := NewDescription(input)

	assert.Error(t, err)
}

func TestDescription_Descriptionが400文字を超えたときエラーになるか(t *testing.T) {
	var input string
	for i := 0; i < 401; i++ {
		input += "a"
	}
	_, err := NewDescription(input)

	assert.Error(t, err)
}

func TestTitle_正常にTitleが設定できるか(t *testing.T) {
	input := "テストタイトル"
	title, err := NewTitle(input)

	assert.NoError(t, err)
	assert.Equal(t, input, title.Value())
}

func TestTitle_Titleが未入力のときエラーになるか(t *testing.T) {
	input := ""
	_, err := NewTitle(input)
	assert.Error(t, err)
}

func TestTitle_Titleが100文字を超えたときエラーになるか(t *testing.T) {
	var input string
	for i := 0; i < 101; i++ {
		input += "a"
	}
	_, err := NewTitle(input)
	assert.Error(t, err)
}
