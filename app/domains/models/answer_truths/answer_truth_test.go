package answer_truths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerTruthId_正常にAnswerTruthIdを設定できるか(t *testing.T) {
	id, err := NewAnswerTruthId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestAnswerTruthId_AnswerTruthIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewAnswerTruthId(0)

	assert.Error(t, err)
}
