package regex

import "regexp"

// 是否为一个邮箱地址.
func IsEmail(str string) bool {
	m, err := regexp.MatchString("", str)
	if err != nil {
		return false
	}
	return m
}

// 是否为一个手机号码.
func IsPhone(str string) bool {
	m, err := regexp.MatchString("", str)
	if err != nil {
		return false
	}
	return m
}