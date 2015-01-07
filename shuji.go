package main

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vimrus/shuji/handlers"
	"net/http"
)

type pongoRender struct {
	cache map[string]*pongo2.Template
}

func newPongoRender() *pongoRender {
	return &pongoRender{map[string]*pongo2.Template{}}
}

func writeHeader(w http.ResponseWriter, code int, contentType string) {
	if code >= 0 {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(code)
	}
}

func (p *pongoRender) Render(w http.ResponseWriter, code int, data ...interface{}) error {
	file := data[0].(string)
	ctx := data[1].(pongo2.Context)
	var t *pongo2.Template

	if tmpl, ok := p.cache[file]; ok {
		t = tmpl
	} else {
		tmpl, err := pongo2.FromFile(file)
		if err != nil {
			return err
		}
		p.cache[file] = tmpl
		t = tmpl
	}
	writeHeader(w, code, "text/html")
	return t.ExecuteWriter(ctx, w)
}

func main() {
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("s3cr3t"))
	r.Use(sessions.Sessions("session", store))

	//r.LoadHTMLTemplates("templates/*")
	r.HTMLRender = newPongoRender()
	r.Static("/static", "static")

	r.GET("/", handlers.Index)
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.Signin)
	r.GET("/logout", handlers.Logout)
	r.GET("/register", handlers.Register)
	r.POST("/register", handlers.Signup)

	r.Run(":8080")
}
