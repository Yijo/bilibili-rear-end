package regex

import (
	"testing"
	"fmt"
)

// 是否为一个昵称正确格式.
func TestIsNickName(t *testing.T) {
	result := IsNickName("")

	if result {
		fmt.Println("这是一个合法的昵称")
	} else {
		fmt.Println("这不是一个合法的昵称")
	}
}