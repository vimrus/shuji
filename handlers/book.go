package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vimrus/shuji/models"
)

func New(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	name := username
	c.HTML(200, "templates/new.html", pongo2.Context{
		"name":     name,
		"username": username,
	})
}

func Books(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	name := username.(string)
	books := models.GetUserBooks(name)
	if cap(books) == 0 {
		c.Redirect(301, "/new")
	} else {
		c.HTML(200, "templates/books.html", pongo2.Context{
			"name":     name,
			"username": username,
			"books":    books,
		})
	}
}
