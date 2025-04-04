package type_converts

import (
	"testing"

	"github.com/google/uuid"
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

func TestStringToUuidOrUuidNil_空文字列の時にNilが返されるか(t *testing.T) {
	value, err := StringToUuidOrNil("")

	assert.Nil(t, value)
	assert.Nil(t, err)
}

func TestStringToUuidOrUuidNil_正常な値の時に変換されているか(t *testing.T) {
	// 有効なUUID文字列
	validUuid := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	expected, _ := uuid.Parse(validUuid)

	value, err := StringToUuidOrNil(validUuid)

	assert.Equal(t, expected, *value)
	assert.Nil(t, err)
}

func TestStringToUuidOrUuidNil_正常じゃない値の時にエラーが返されるか(t *testing.T) {
	// 無効なUUID文字列
	invalidUuid := "not-a-uuid-string"

	value, err := StringToUuidOrNil(invalidUuid)

	assert.NotNil(t, err)
	assert.Nil(t, value) // 無効な場合はゼロ値のUUIDが返される
}

func TestPointerUuidToString_正常な値の時に変換されるか(t *testing.T) {
	uuid := uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

	value := PointerUuidToString(&uuid)

	assert.Equal(t, uuid.String(), value)
}

func TestPointerUuidToString_nilの時に空文字に変換されるか(t *testing.T) {

	value := PointerUuidToString(nil)

	assert.Equal(t, "", value)
}
