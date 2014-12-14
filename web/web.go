package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type (
	HandlerFunc func(*Context)
	H           map[string]interface{}
	Handler     struct {
		patt   string
		Handle HandlerFunc
	}
	Engine struct {
		pool     sync.Pool
		handlers map[string][]*Handler
	}
)

func New() *Engine {
	engine := &Engine{}
	engine.handlers = make(map[string][]*Handler)
	engine.pool.New = func() interface{} {
		ctx := &Context{}
		ctx.Writer = &ctx.writer
		return ctx
	}
	return engine
}

func (engine *Engine) add(meth string, handler *Handler) {
	engine.handlers[meth] = append(engine.handlers[meth], handler)
}

func (engine *Engine) HEAD(patt string, handler HandlerFunc) {
	engine.add("HEAD", &Handler{patt, handler})
}

func (engine *Engine) GET(patt string, handler HandlerFunc) {
	engine.add("GET", &Handler{patt, handler})
	engine.add("HEAD", &Handler{patt, handler})
}

func (engine *Engine) POST(patt string, handler HandlerFunc) {
	engine.add("POST", &Handler{patt, handler})
}

func (engine *Engine) PUT(patt string, handler HandlerFunc) {
	engine.add("PUT", &Handler{patt, handler})
}

func (engine *Engine) DELETE(patt string, handler HandlerFunc) {
	engine.add("DELETE", &Handler{patt, handler})
}

func (engine *Engine) OPTIONS(patt string, handler HandlerFunc) {
	engine.add("OPTIONS", &Handler{patt, handler})
}

func (engine *Engine) PATCH(patt string, handler HandlerFunc) {
	engine.add("PATCH", &Handler{patt, handler})
}

//https://github.com/daryl/zeus
//Copyright (c) 2014 Daryl Ginn
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > 8 && r.URL.Path[:8] == "/static/" {
		fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
		fileServer.ServeHTTP(w, r)
		return
	}

	plen := len(r.URL.Path)
	// Redirect trailing slash URL's.
	if plen > 1 && r.URL.Path[plen-1:] == "/" {
		http.Redirect(w, r, r.URL.Path[:plen-1], 301)
		return
	}

	c := engine.createContext(w, r)

	// Map over the registered handlers for
	// the current request (if there is any).
	for _, handler := range engine.handlers[r.Method] {

		// Try the pattern against the URL path.
		if vars, ok := handler.try(r.URL.Path); ok {
			// Prepend params to URL query.
			r.URL.RawQuery = vars.Encode() + "&" + r.URL.RawQuery
			// Serve handlers.
			handler.Handle(c)
			return
		}
	}

	// Default 404.
	http.NotFound(w, r)
}

func (h *Handler) try(path string) (url.Values, bool) {
	// Patt and URL segments.
	ps := strings.Split(h.patt[1:], "/")
	us := strings.Split(path[1:], "/")

	// If the patt and URL slices
	// have different lengths we
	// already know it's bad.
	if len(ps) != len(us) {
		return nil, false
	}

	// Compiled.
	var cs string
	// Parameters.
	uv := url.Values{}

	for idx, part := range ps {
		// Part is at least :x
		if len(part) > 1 && part[:1] == ":" {
			// Add to parameters.
			uv.Add(part[1:], us[idx])
			// Add URL seg.
			cs += "/" + us[idx]
			continue
		}
		// Add patt seg.
		cs += "/" + part
	}
	// Match?
	if cs == path {
		return uv, true
	}

	return nil, false
}

func (engine *Engine) Listen(addr string) {
	fmt.Printf("Listening and serving HTTP on %s\n", addr)
	http.ListenAndServe(addr, engine)
}
