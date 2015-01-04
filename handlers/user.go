package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	//"github.com/vimrus/shuji/models"
)

type RegisterJson struct {
	Email    string `json:"email" binding:"required"`
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	c.HTML(200, "templates/login.html", pongo2.Context{})
}

func Signin(c *gin.Context) {
}

func Logout(c *gin.Context) {
}

func Register(c *gin.Context) {
	c.HTML(200, "templates/register.html", pongo2.Context{})
}

func Signup(c *gin.Context) {
	var json RegisterJson
	c.Bind(&json)

	c.JSON(200, gin.H{"Ok": "true"})
}
