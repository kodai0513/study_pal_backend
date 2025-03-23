package entities

import (
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/master_datas/master_roles"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkbook_正常に値が設定できるか(t *testing.T) {
	input := struct {
		description string
		title       string
	}{
		description: "プログラミングの基礎を確認する問題集です",
		title:       "プログラミング基礎問題集",
	}
	description, _ := workbooks.NewDescription(input.description)
	title, _ := workbooks.NewTitle(input.title)
	userId := users.CreateUserId()
	workbookId := workbooks.CreateWorkbookId()
	workbook := CreateWorkbook(workbookId, description, userId, title)

	assert.Equal(t, input.description, workbook.Description())
	assert.Equal(t, false, workbook.IsPublic())
	assert.Equal(t, input.title, workbook.Title())
	assert.Equal(t, userId.Value(), workbook.UserId())
	assert.Equal(t, workbookId.Value(), workbook.Id())
	assert.Equal(t, 1, len(workbook.WorkbookMembers()))
	assert.Equal(t, master_roles.Admin, workbook.WorkbookMembers()[0].RoleId())
	assert.Equal(t, userId.Value(), workbook.WorkbookMembers()[0].UserId())
	assert.Equal(t, workbookId.Value(), workbook.WorkbookMembers()[0].WorkbookId())
}
