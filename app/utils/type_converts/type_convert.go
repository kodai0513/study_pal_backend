package type_converts

import "strconv"

func StringToInt(value string, defaultValue int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return result
}
