package handlers

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"fmt"
)

// 登录账号.
func SignInHandler(ctx *gin.Context) {
	// 用户名
	username := ctx.PostForm("username")

	fmt.Println(username)
}


// 登出账号.
func SignOutHandler(ctx *gin.Context) {
	// 用户名
	username := ctx.PostForm("username")
	fmt.Println(username)
}