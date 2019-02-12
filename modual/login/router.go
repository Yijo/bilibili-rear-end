package login

import (
	"github.com/gin-gonic/gin"
	"bilibili-rear-end/modual/login/handlers"
)

func SubRouters(r *gin.Engine) {
	router := r.Group("")
	{
		// 登录
		router.POST("/login", handlers.SignInHandler)

		// 邮箱注册
		router.POST("/register/mail", handlers.RegisterMailHandler)

		// 手机注册
		router.POST("/register/phone", handlers.RegisterPhoneHandler)

		// 获取注册验证码.
		router.GET("/register/captcha", handlers.GetRegisterCaptchaHandler)
	}
}