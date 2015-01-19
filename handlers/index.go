package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/vimrus/gin"
	"github.com/vimrus/shuji/utils/sessions"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	account := session.Get("account")
	if account == nil {
		c.HTML(200, "index.html", pongo2.Context{})
	} else {
		c.HTML(200, "home.html", pongo2.Context{
			"account": account,
		})
	}
}
