package handlers

import (
	"testing"
	"fmt"
)

// 登录.
func TestSignInHandler(t *testing.T) {

}

// 登出.
func TestSignOutHandler(t *testing.T) {

}

func Test1(t *testing.T) {
	map1 := make(map[string]string)
	//map2
	fmt.Printf("%p\n", &map1)
	mapTo(map1)
	fmt.Printf("%p\n",&map1)
	}

func mapTo(map1 map[string]string) {

	map1["key"] = "key"
	fmt.Printf("%p\n",&map1)
}