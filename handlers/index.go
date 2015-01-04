package handlers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"github.com/vimrus/shuji/models"
)

func Index(c *gin.Context) {
	name := models.GetUserByName("test")
	session := sessions.Get(c)
	session.Set("hello", "world")
	session.Save()
	name = session.Get("hello").(string)

	c.HTML(200, "templates/index.html", pongo2.Context{"name": name})
}
