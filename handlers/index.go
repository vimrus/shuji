package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/vimrus/gin"
	"github.com/vimrus/shuji/utils/sessions"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get("account")
	name := account
	if name == nil {
		c.HTML(200, "templates/index.html", pongo2.Context{})
	} else {
		c.HTML(200, "templates/home.html", pongo2.Context{
			"name":    name,
			"account": account,
		})
	}
}
