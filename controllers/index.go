package controllers

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl, err := pongo2.FromString("Hello {{ name|capfirst }}!")
	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(pongo2.Context{"name": "shuji"})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, out)
}
