package util

import "strconv"

func ConvertToInteger(String string) int {
	result, _ := strconv.Atoi(String)
	return result
}
