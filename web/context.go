package web

import (
	"github.com/flosch/pongo2"
	"net/http"
)

type Context struct {
	Request *http.Request
	writer  responseWriter
	Writer  ResponseWriter
	Data    map[string]interface{}
}

func (engine *Engine) createContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := engine.pool.Get().(*Context)
	ctx.Request = r
	ctx.Data = nil
	ctx.writer = responseWriter{
		ResponseWriter: w,
		status:         200,
		size:           NoWritten,
	}
	return ctx
}

func (ctx *Context) HTML(code int, file string, obj map[string]interface{}) {

	tpl := pongo2.Must(pongo2.FromFile("templates/" + file))
	err := tpl.ExecuteWriter(obj, ctx.Writer)

	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}
