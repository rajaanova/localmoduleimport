package utils

import "strconv"

func GetConversion(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
