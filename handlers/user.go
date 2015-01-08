package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/vimrus/shuji/models"
)

type LoginForm struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterForm struct {
	Email    string `form:"email" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	c.HTML(200, "templates/login.html", pongo2.Context{})
}

func Signin(c *gin.Context) {
	var form LoginForm
	c.BindWith(&form, binding.Form)

	user, err := models.Authorize(form.Name, form.Password)
	if err != nil {
		c.JSON(200, gin.H{"status": "fail", "error": err})
	} else {
		session := sessions.Default(c)
		session.Set("username", user.Name)
		session.Set("avatar", user.Avatar)
		session.Save()

		c.JSON(200, gin.H{"Ok": "true"})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()
	c.Redirect(301, "/")
}

func Register(c *gin.Context) {
	c.HTML(200, "templates/register.html", pongo2.Context{})
}

func Signup(c *gin.Context) {
	var form RegisterForm
	c.BindWith(&form, binding.Form)

	user, err := models.CreateUser(form.Name, form.Email, form.Password)
	if err != nil {
		panic(err)
	}
	session := sessions.Default(c)
	session.Set("username", user.Name)
	session.Set("avatar", user.Avatar)
	session.Save()

	c.JSON(200, gin.H{"Ok": "true"})
}
