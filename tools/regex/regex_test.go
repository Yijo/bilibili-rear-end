package regex

import (
	"testing"
	"fmt"
)

// 是否为一个邮箱地址.
func TestIsEmail(t *testing.T) {
	result := IsEmail("developer.yijo@gmail.com")

	if result {
		fmt.Println("这是一个邮箱地址")
	} else {
		fmt.Println("这不是一个邮箱地址")
	}
}

// 是否为一个手机号码.
func TestIsPhone(t *testing.T) {
	result := IsPhone("")

	if result {
		fmt.Println("这是一个手机号码")
	} else {
		fmt.Println("这不是一个手机号码")
	}
}