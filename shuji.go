package main

import (
	"github.com/vimrus/gin"
	"github.com/vimrus/shuji/handlers"
	"github.com/vimrus/shuji/utils"
	"github.com/vimrus/shuji/utils/sessions"
)

func main() {
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("s3cr3t"))
	r.Use(sessions.Sessions("session", store))

	//r.LoadHTMLTemplates("templates/*")
	r.HTMLRender = utils.NewPongoRender()
	r.Static("/static", "static")

	r.GET("/", handlers.Index)
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.Signin)
	r.GET("/logout", handlers.Logout)
	r.GET("/register", handlers.Register)
	r.POST("/register", handlers.Signup)
	r.GET("/new", handlers.New)
	r.POST("/new", handlers.Create)
	r.GET("/books", handlers.Books)
	r.GET("/book/:book", handlers.Book)

	r.Run(":8080")
}
