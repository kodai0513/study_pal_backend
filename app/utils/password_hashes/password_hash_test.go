package password_hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToHashPassword_正常にハッシュパスワードが生成できるか(t *testing.T) {
	hashPassword, err := ConvertToHashPassword("password")
	assert.NoError(t, err)

	if len(hashPassword) <= 0 {
		t.Fatal("empty hashPassword")
	}
}

func TestCheckPasswordHash_正常にパスワードチェックができるか(t *testing.T) {
	inputPassword := "password"
	hashPassword, _ := ConvertToHashPassword(inputPassword)
	assert.NoError(t, CheckPasswordHash(inputPassword, hashPassword))
}
