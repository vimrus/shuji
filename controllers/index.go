package controllers

import (
	"github.com/vimrus/shuji/web"
)

func Index(c *web.Context) {
	c.HTML(200, "index.html", web.H{"name": "Shuji"})
}
