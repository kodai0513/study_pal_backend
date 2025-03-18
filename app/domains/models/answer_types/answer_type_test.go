package answer_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerTypeId_正常にAnswerTypeIdを設定できるか(t *testing.T) {
	id, err := NewAnswerTypeId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestAnswerTypeId_AnswerTypeIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewAnswerTypeId(0)

	assert.Error(t, err)
}
