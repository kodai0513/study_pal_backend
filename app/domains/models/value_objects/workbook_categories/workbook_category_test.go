package workbook_categories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName_正常Nameが設定できるか(t *testing.T) {
	input := "テスト科目"
	name, err := NewName(input)

	assert.NoError(t, err)
	assert.Equal(t, input, name.Value())
}

func TestName_Nameが未入力のときエラーになるか(t *testing.T) {
	input := ""
	_, err := NewName(input)

	assert.Error(t, err)
}

func TestName_Nameが30文字を超えるときエラーになるか(t *testing.T) {
	var input string
	for i := 0; i < 31; i++ {
		input += "a"
	}
	_, err := NewName(input)

	assert.Error(t, err)
}
