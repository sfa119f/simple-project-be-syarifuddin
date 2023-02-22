package utils

import (
	"fmt"
	"strings"
)

func ArrIntToStr(arr []int64) string {
	return strings.Trim(
		strings.Join(strings.Fields(fmt.Sprint(arr)), ", "), "[]",
	)
}

func ArrContainsStr(arr []string, val string) bool {
	for _, el := range arr {
		if el == val { return true }
	}
	return false
}
