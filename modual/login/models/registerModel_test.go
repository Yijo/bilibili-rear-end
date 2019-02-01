package models

import (
	"testing"
	"fmt"
)

// 查询用户登录所需信息.
func TestQueryUser(t *testing.T) {
	user, err := QueryUser("")
	if err != nil {
		fmt.Println("出错了，错误信息:", err)
	} else {
		fmt.Println("用户信息:", user)
	}
}