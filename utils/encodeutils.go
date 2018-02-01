package base62

import (
	"math"
	"strings"
)

const encodeBase = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const encodeBits = 62

//Function that accepts a Base10 value and encodes to Base64
func ToBase62(input int) string{
	var result = ""

	//-- First Pass ----------
	var factor = input / encodeBits
	var remainder = input % encodeBits
	result = result + string(encodeBase[remainder])

	//-- Process Remaining Bits ----------
	var leftover = int(math.Floor(float64(factor)))

	for leftover != 0 {
		var remainder = leftover % encodeBits
		var factor = leftover / encodeBits
		result = string(encodeBase[int(remainder)]) + result

		leftover = int(math.Floor(float64(factor)))
	}

	return string(result)
}

//Function that accepts a Base64 value and encodes to Base10
func ToBase10(str string) int{
	var result = 0

	for _, value := range str {
		result = (encodeBits * result) + strings.Index(encodeBase, string(value))
	}
	return result
}