package roles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleId_正常にRoleIdを設定できるか(t *testing.T) {
	id, err := NewRoleId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestRoleId_RoleIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewRoleId(0)

	assert.Error(t, err)
}
