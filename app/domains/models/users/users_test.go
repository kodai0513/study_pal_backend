package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserId_正常なUserId(t *testing.T) {
	input := 1
	userId, err := NewUserId(input)

	assert.NoError(t, err)
	assert.Equal(t, 1, userId.Value())
}

func TestUserId_異常なUserId(t *testing.T) {
	input := 0
	userId, err := NewUserId(input)

	assert.Error(t, err)
	assert.Equal(t, 0, userId.Value())
}

func TestNewEmail_正常なメールアドレス(t *testing.T) {
	input := "test@example.com"
	email, err := NewEmail(input)

	assert.NoError(t, err)
	assert.Equal(t, input, email.Value())
}

func TestNewEmail_空のメールアドレスはエラー(t *testing.T) {
	input := ""
	email, err := NewEmail(input)

	assert.Error(t, err)
	assert.Empty(t, email.Value())
}

func TestNewEmail_無効なメールアドレスはエラー(t *testing.T) {
	invalidEmails := []string{
		"plainaddress",          // ドメインなし
		"@missing-username.com", // ユーザー名なし
		"missing-at.com",        // @がない
		"test@.com",             // ドメインが不正
		"test@com",              // 不完全なドメイン
		"test@domain..com",      // 連続したドット
	}

	for _, input := range invalidEmails {
		t.Run("Invalid Email: "+input, func(t *testing.T) {
			email, err := NewEmail(input)

			assert.Error(t, err)
			assert.Empty(t, email.Value())
		})
	}
}

func TestNewName_正常な値の時(t *testing.T) {
	input := "JohnDoe"
	name, err := NewName(input)

	assert.NoError(t, err)
	assert.Equal(t, input, name.Value())
}

func TestNewName_名前が空の時エラー(t *testing.T) {
	input := ""
	name, err := NewName(input)

	assert.Error(t, err)
	assert.Empty(t, name.Value())
}

func TestNewName_名前が英数字でない時エラー(t *testing.T) {
	input := "名前123!"
	name, err := NewName(input)

	assert.Error(t, err)
	assert.Empty(t, name.Value())
}

func TestNewName_名前が20文字を超えた時エラー(t *testing.T) {
	input := "korehanamaegahijouninagaidesu123"
	name, err := NewName(input)

	assert.Error(t, err)
	assert.Empty(t, name.Value())
}

func TestNewNickName_正常な値の時(t *testing.T) {
	input := "ValidNickName"
	nickName, err := NewNickName(input)

	assert.NoError(t, err)
	assert.Equal(t, input, nickName.Value())
}

func TestNewNickName_名前が20文字を超えた時エラー(t *testing.T) {
	input := "ThisNickNameIsWayTooLong"
	nickName, err := NewNickName(input)

	assert.Error(t, err)
	assert.Empty(t, nickName.Value())
}

func TestNewPassword_正常な値の時(t *testing.T) {
	input := "ValidPassword123"
	password, err := NewHashPassword(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, password.Value()) // ハッシュ化されたパスワードが空でないことを確認
}

func TestNewPassword_パスワードが空の時エラー(t *testing.T) {
	input := ""
	password, err := NewHashPassword(input)

	assert.Error(t, err)
	assert.Empty(t, password.Value())
}
