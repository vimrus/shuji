package controllers

import (
	"github.com/vimrus/shuji/web"
)

func Login(c *web.Context) {
	c.HTML(200, "login.html", web.H{})
}

func Logout(c *web.Context) {
}

func Signup(c *web.Context) {
	c.HTML(200, "signup.html", web.H{})
}
