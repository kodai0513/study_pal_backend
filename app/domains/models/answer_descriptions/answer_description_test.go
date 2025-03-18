package answer_descriptions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerDescriptionId_正常にAnswerDescriptionIdを設定できるか(t *testing.T) {
	id, err := NewAnswerDescriptionId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestAnswerDescriptionId_AnswerDescriptionIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewAnswerDescriptionId(0)

	assert.Error(t, err)
}

func TestName_正常にNameが設定できるか(t *testing.T) {
	input := "テスト"
	name, err := NewName(input)

	assert.NoError(t, err)
	assert.Equal(t, input, name.Value())
}

func TestName_Nameが未入力のときエラーになるか(t *testing.T) {
	input := ""
	_, err := NewName(input)

	assert.Error(t, err)
}

func TestName_Nameが100文字を超えたときエラーになるか(t *testing.T) {
	var input string
	for i := 0; i < 101; i++ {
		input += "a"
	}
	_, err := NewName(input)

	assert.Error(t, err)
}
