package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/master_datas/master_roles"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWorkbook_正常に登録できるか(t *testing.T) {
	input := struct {
		description string
		title       string
	}{
		description: "プログラミングの基礎を確認する問題集です",
		title:       "プログラミング基礎問題集",
	}
	description, _ := workbooks.NewDescription(input.description)
	title, _ := workbooks.NewTitle(input.title)
	userId := uuid.New()
	workbookId := uuid.New()
	workbook := CreateWorkbook(workbookId, description, userId, title)

	assert.Equal(t, input.description, workbook.Description())
	assert.Equal(t, false, workbook.IsPublic())
	assert.Equal(t, input.title, workbook.Title())
	assert.Equal(t, userId, workbook.UserId())
	assert.Equal(t, workbookId, workbook.Id())
	assert.Equal(t, false, workbook.IsPublic())
}

func TestWorkbook_正常に更新できるか(t *testing.T) {
	description, err := workbooks.NewDescription("テスト説明")
	assert.NoError(t, err)
	title, err := workbooks.NewTitle("テスト")
	assert.NoError(t, err)
	userId := uuid.New()
	workbookId := uuid.New()
	workbookMembers := []*WorkbookMember{NewWorkbookMember(uuid.New(), master_roles.Admin, userId, workbookId)}
	workbook := NewWorkbook(
		workbookId,
		description,
		[]uuid.UUID{uuid.New(), uuid.New()},
		false,
		[]uuid.UUID{uuid.New(), uuid.New()},
		title,
		make([]uuid.UUID, 0),
		userId,
		workbookMembers,
	)

	updateDescription, err := workbooks.NewDescription("テスト説明変更")
	assert.NoError(t, err)
	updateTitle, err := workbooks.NewTitle("テスト変更")
	assert.NoError(t, err)
	workbook.SetDescription(updateDescription)
	workbook.SetTitle(updateTitle)
	err = workbook.ChangePublic()
	assert.NoError(t, err)

	assert.Equal(t, updateDescription.Value(), workbook.Description())
	assert.Equal(t, updateTitle.Value(), workbook.Title())
	assert.Equal(t, true, workbook.IsPublic())

	workbook.ChangePrivate()
	assert.Equal(t, false, workbook.IsPublic())
}

func TestWorkbook_問題が無い時に公開できないようになっているか(t *testing.T) {
	description, err := workbooks.NewDescription("テスト説明")
	assert.NoError(t, err)
	title, err := workbooks.NewTitle("テスト")
	assert.NoError(t, err)
	userId := uuid.New()
	workbookId := uuid.New()
	workbookMembers := []*WorkbookMember{NewWorkbookMember(uuid.New(), master_roles.Admin, userId, workbookId)}
	workbook := NewWorkbook(
		workbookId,
		description,
		make([]uuid.UUID, 0),
		false,
		make([]uuid.UUID, 0),
		title,
		make([]uuid.UUID, 0),
		userId,
		workbookMembers,
	)

	err = workbook.ChangePublic()
	assert.Error(t, err)
	assert.Equal(t, false, workbook.IsPublic())
}
