package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatement_正常にStatementが設定できるか(t *testing.T) {
	input := "テスト"
	statement, err := NewStatement(input)

	assert.NoError(t, err)
	assert.Equal(t, input, statement.Value())
}

func TestStatement_Statementが未入力のときエラーになるか(t *testing.T) {
	input := ""
	_, err := NewStatement(input)

	assert.Error(t, err)
}

func TestStatement_Statementが1000文字を超えたときエラーになるか(t *testing.T) {
	var input string
	for i := 0; i < 1001; i++ {
		input += "a"
	}
	_, err := NewStatement(input)

	assert.Error(t, err)
}
