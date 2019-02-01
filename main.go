package main

import (
	"github.com/gin-gonic/gin"
	"bilibili-rear-end/modual/login"
	"fmt"
	"bilibili-rear-end/middleware"
	"bilibili-rear-end/configer"
	//"bilibili-rear-end/database/mysql"
)

func main() {
	configer.InitConfig() // Init config


	router := gin.Default()

	router.Use(middleware.MiddleWare())	// open global middleware, use with routes

	//router.BasePath() = "/v1"


	login.SubRouters(router) // load routes

	fmt.Println(router.RouterGroup)

	router.Run()	// start routes, default port 8080

	//mysql.OpenDB()

	//defer mysql.

}
