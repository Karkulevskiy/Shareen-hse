package utils

import (
	"math/rand"
	"slices"
	"strings"
	"unicode/utf8"
)

func CreateURL() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var alpabetLen = uint32(utf8.RuneCountInString(alphabet))

	var (
		nums    []uint32
		num     = rand.Uint32()
		builder strings.Builder
	)
	for num > 0 {
		nums = append(nums, num%alpabetLen)
		num /= alpabetLen
	}
	slices.Reverse(nums)
	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}
