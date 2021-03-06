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

// Register an account with an email.
func RegisterMailHandler(ctx *gin.Context) {

	registerAccount(ctx, emailStyle)

}

// Register an account with an phone.
func RegisterPhoneHandler(ctx *gin.Context) {
	registerAccount(ctx, phoneStyle)
}
 
// Register an account.
func registerAccount(ctx *gin.Context, style registerStyle) {

	// validate nickname
	nickName := ctx.PostForm("nickname")
	if nickName == "" {
		network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("nickname is nil.").Response(ctx)
	}

	result := regex.IsNickName(nickName)
	if !result {
		network.Failure(http.StatusInternalServerError, nickNameWrongFormate).Response(ctx)
	}


	// password
	password := ctx.PostForm("password")
	if password == "" {
		network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("password is nil.").Response(ctx)
	}


	// account
	var account string
	switch style {
	case phoneStyle:
		// validate phone
		phone := ctx.PostForm("phone")
		if phone == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("phone is nil.").Response(ctx)
		}

		result := regex.IsPhone(phone)
		if !result {
			network.Failure(http.StatusBadRequest, phoneWrongFormate).Response(ctx)
		}

		// phone captcha
		phoneVerificationCode := ctx.PostForm("phone_verification_code")
		if phoneVerificationCode == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("phone_verification_code is nil.").Response(ctx)
		}

	case emailStyle:
		// validate email address
		email := ctx.PostForm("email")
		if email == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("email is nil.").Response(ctx)
		}

		result := regex.IsEmail(email)
		if !result {
			network.Failure(http.StatusBadRequest, emailWrongFormate).Response(ctx)
		}

		// email address captcha
		emailVerificationCode := ctx.PostForm("email_verification_code")
		if emailVerificationCode == "" {
			network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("email_verification_code is nil.").Response(ctx)
		}
	}


	// 查询用户名是否被注册
	act, err := models.QueryUser(account)
	if err != nil {
		network.Failure(http.StatusInternalServerError, registerFailure).AppendMessage(err.Error()).Response(ctx)
	}

	if act != "" {

	}

	network.Success(http.StatusOK, network.SUCCESS, none, "").Response(ctx)
}