package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/vimrus/shuji/models"
)

type CreateForm struct {
	Name string `form:"name" binding:"required"`
	Desc string `form:"desc"`
}

func New(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get("account")
	name := account
	c.HTML(200, "templates/new.html", pongo2.Context{
		"name":    name,
		"account": account,
	})
}

func Create(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get("userid").(int64)

	var form CreateForm
	c.BindWith(&form, binding.Form)

	_, err := models.CreateBook(form.Name, form.Desc, userid)
	if err != nil {
		c.JSON(200, gin.H{"status": "fail", "error": err})
	} else {
		c.JSON(200, gin.H{"Ok": "true", "bookName": form.Name})
	}
}

func Books(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get("account")
	name := account.(string)
	books := models.GetUserBooks(name)
	if cap(books) == 0 {
		c.Redirect(301, "/new")
	} else {
		c.HTML(200, "templates/books.html", pongo2.Context{
			"name":    name,
			"account": account,
			"books":   books,
		})
	}
}
