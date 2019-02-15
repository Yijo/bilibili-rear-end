package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"bilibili-rear-end/tools/regex"
	"bilibili-rear-end/network"
	"bilibili-rear-end/tools/captcha"
	"bilibili-rear-end/database/nosql/redis"
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
	default:
		network.Failure(http.StatusBadRequest, parameterFailure).AppendMessage("undefined style").Response(ctx)
	}

	capt, err := captcha.GetCaptcha(style, 2)
	if err != nil {
		network.Failure(http.StatusInternalServerError, failureSendCaptcha).Response(ctx)
	}

	redis.GetInstance().Set("", capt)

	network.Success(http.StatusOK, getCaptchaSuccess, nil).Response(ctx)
}
