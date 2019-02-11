package main

import (
	"github.com/gin-gonic/gin"
	"bilibili-rear-end/modual/login"
	"fmt"
	"bilibili-rear-end/middleware"
	"bilibili-rear-end/configer"
	"net/http"
	"time"
)

func main() {
	configer.InitConfig() // Init config

	router := gin.Default()
	severListen(router)	// start server listen

	router.Use(middleware.MiddleWare())	// open global middleware, use with routes

	//router.BasePath() = "/v1"


	login.SubRouters(router) // load routes

	fmt.Println(router.RouterGroup)

	router.Run()	// start routes, default port 8080

}

// Server listen.
func severListen(router *gin.Engine) {
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
