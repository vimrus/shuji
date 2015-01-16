package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/vimrus/gin"
	"github.com/vimrus/gin/binding"
	"github.com/vimrus/shuji/models"
	"github.com/vimrus/shuji/utils/sessions"
)

type LoginForm struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterForm struct {
	Email    string `form:"email" binding:"required"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	c.HTML(200, "login.html", pongo2.Context{})
}

func Signin(c *gin.Context) {
	var form LoginForm
	c.BindWith(&form, binding.Form)

	user, err := models.Authorize(form.Account, form.Password)
	if err != nil {
		c.JSON(200, gin.H{"status": "fail", "error": err})
	} else {
		session := sessions.Default(c)
		session.Set("userid", user.Id)
		session.Set("account", user.Account)
		session.Set("avatar", user.Avatar)
		session.Save()

		c.JSON(200, gin.H{"Ok": "true"})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("account")
	session.Save()
	c.Redirect(301, "/")
}

func Register(c *gin.Context) {
	c.HTML(200, "register.html", pongo2.Context{})
}

func Signup(c *gin.Context) {
	var form RegisterForm
	c.BindWith(&form, binding.Form)

	user, err := models.CreateUser(form.Account, form.Email, form.Password)
	if err != nil {
		panic(err)
	}
	session := sessions.Default(c)
	session.Set("userid", user.Id)
	session.Set("account", user.Account)
	session.Set("avatar", user.Avatar)
	session.Save()

	c.JSON(200, gin.H{"Ok": "true"})
}
