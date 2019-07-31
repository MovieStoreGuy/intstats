package utils

import "strconv"

// ParseInt will accept a string and parse it
// if the value can't be parsed, the value will be zero
func ParseInt(value string) uint64 {
	val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		val = 0
	}
	return val
}
