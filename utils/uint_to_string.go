package utils

import (
	"strconv"
)

func UintToString(num uint) string {
	str := strconv.FormatUint(uint64(num), 10)
	return str
}
