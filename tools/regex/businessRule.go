package regex

import "regexp"

// 是否为一个昵称正确格式.
func IsNickName(str string) bool {
	m, err := regexp.MatchString("", str)
	if err != nil {
		return false
	}
	return m
}