package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"bilibili-rear-end/tools/regex"
	"bilibili-rear-end/network"
	"bilibili-rear-end/modual/login/models"
)

type registerStyle int
const (
	phoneStyle = iota
	emailStyle
)

// 注册账号(邮箱).
func RegisterMailHandler(ctx *gin.Context) {

	registerAccount(ctx, emailStyle)

}

// 注册账号(手机).
func RegisterPhoneHandler(ctx *gin.Context) {
	registerAccount(ctx, phoneStyle)
}
 
// 通用注册(根据注册方式不同注册).
func registerAccount(ctx *gin.Context, style registerStyle) {
	// 昵称
	nickName := ctx.PostForm("nickname")
	if nickName == "" {
		network.Failure(http.StatusBadRequest, parameterFailure).AppendError("nickname is nil.").Response(ctx)
	}

	// 验证昵称合法性
	result := regex.IsNickName(nickName)
	if !result {
		network.Success(http.StatusOK, network.FAILED, nickNameWrongFormate, nil).Response(ctx)
	}

	// 密码
	password := ctx.PostForm("password")
	if password == "" {
		network.Failure(http.StatusBadRequest, parameterFailure).AppendError("password is nil.").Response(ctx)
	}

	var account string = ""
	switch style {
	case phoneStyle:
		// 手机
		phone := ctx.PostForm("phone")
		if phone == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendError("phone is nil.").Response(ctx)
		}

		// 验证手机合法性
		result := regex.IsPhone(phone)
		if !result {
			network.Success(http.StatusBadRequest, network.FAILED, phoneWrongFormate, nil).Response(ctx)
		}

		// 手机验证码
		phoneVerificationCode := ctx.PostForm("phone_verification_code")
		if phoneVerificationCode == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendError("phone_verification_code is nil.").Response(ctx)
		}
	case emailStyle:
		// 邮箱
		email := ctx.PostForm("email")
		if email == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendError("email is nil.").Response(ctx)
		}

		// 验证邮箱合法性
		result := regex.IsEmail(email)
		if !result {
			network.Success(http.StatusBadRequest, network.FAILED, emailWrongFormate, nil).Response(ctx)
		}

		// 邮箱验证码
		emailVerificationCode := ctx.PostForm("email_verification_code")
		if emailVerificationCode == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendError("email_verification_code is nil.").Response(ctx)
		}
	}

	// 查询用户名是否被注册
	act, err := models.QueryUser(account)
	if err != nil {
		network.Failure(http.StatusInternalServerError, registerFailure).AppendError(err.Error()).Response(ctx)
	}

	if act != "" {

	}

	network.Success(http.StatusOK, network.SUCCESS, none, "").Response(ctx)
}