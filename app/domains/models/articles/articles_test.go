package articles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticleId_正常にArticleIdを設定できるか(t *testing.T) {
	id, err := NewArticleId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestArticleId_ArticleIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewArticleId(0)

	assert.Error(t, err)
}

func TestNewDescription_正常な値の時(t *testing.T) {
	input := "テスト投稿"
	desc, err := NewDescription(input)

	assert.NoError(t, err)

	assert.Equal(t, input, desc.Value())
}

func TestNewDescription_何を入力されていないときエラー(t *testing.T) {
	input := ""
	desc, err := NewDescription(input)

	assert.Error(t, err)
	assert.Empty(t, desc.Value())
}

func TestNewDescription_文字数が401を超えたときエラー(t *testing.T) {
	input := `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	aaaaaaaaa`

	desc, err := NewDescription(input)

	assert.Error(t, err)
	assert.Empty(t, desc.Value())
}
