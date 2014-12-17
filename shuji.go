package main

import (
	"github.com/vimrus/shuji/controllers"
	"github.com/vimrus/shuji/web"
)

func main() {
	web := web.New()
	web.GET("/", controllers.Index)
	web.GET("/login", controllers.Login)
	web.GET("/logout", controllers.Logout)
	web.GET("/signup", controllers.Signup)

	web.Listen(":8080")
}
