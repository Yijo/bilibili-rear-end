package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"bilibili-rear-end/tools/regex"
	"bilibili-rear-end/network"
)

// Get register captcha.
func GetRegisterCaptchaHandler(ctx *gin.Context) {
	style := ctx.Query("style")
	if style == "" {
		network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("style is nil.").Response(ctx)
	}

}

// Get captcha.
func GetCaptcha(ctx *gin.Context, style int) {

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
