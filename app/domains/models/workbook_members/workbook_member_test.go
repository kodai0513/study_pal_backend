package workbook_members

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkbookMemberId_正常にWorkbookMemberIdを設定できるか(t *testing.T) {
	id, err := NewWorkbookMemberId(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, id.Value())
}

func TestWorkbookMemberId_WorkbookMemberIdが0以下のときはエラーになるか(t *testing.T) {
	_, err := NewWorkbookMemberId(0)

	assert.Error(t, err)
}
