package type_converts

import (
	"strconv"

	"github.com/google/uuid"
)

func StringToInt(value string, defaultValue int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}

func StringToUuidOrNil(value string) (*uuid.UUID, error) {
	if len(value) == 0 {
		return nil, nil
	}

	uuid, err := uuid.Parse(value)

	if err != nil {
		return nil, err
	}
	return &uuid, nil
}

func PointerUuidToString(value *uuid.UUID) string {
	if value == nil {
		return ""
	}

	return value.String()
}
