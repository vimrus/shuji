package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	name := session.Get("username")
	if name == nil {
		name = "shuji"
	}
	c.HTML(200, "templates/index.html", pongo2.Context{"name": name})
}
