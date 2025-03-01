package converts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToInt_正常な値の時に変換されているか(t *testing.T) {
	value := StringToInt("1", 0)

	assert.Equal(t, 1, value)
}

func TestStringToInt_正常じゃない値の時にデフォルト値になるか(t *testing.T) {
	value := StringToInt("test", 0)

	assert.Equal(t, 0, value)
}
