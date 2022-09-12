package util

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func StringToInt(val string, defaultVal int) (newVal int) {
	newVal, err := strconv.Atoi(val)

	if err != nil {
		return defaultVal
	}

	return newVal
}

func StringToInt64(val string, defaultVal int64) (newVal int64) {
	newVal, err := strconv.ParseInt(val, 10, 64)

	if err != nil {
		return defaultVal
	}

	return newVal
}

func ToString(val interface{}, defaultVal string) (newVal string) {
	newVal, ok := val.(string)

	if ok {
		return
	}

	return defaultVal
}

func MD5Cypto(str string) string {
	bytes := []byte(str)

	return fmt.Sprintf("%x", md5.Sum(bytes))
}
