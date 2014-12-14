package main

import (
	"github.com/vimrus/shuji/controllers"
	"github.com/vimrus/shuji/web"
)

func main() {
	web := web.New()
	web.GET("/", controllers.Index)

	web.Listen(":8080")
}
