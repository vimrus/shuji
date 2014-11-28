package controllers

import (
	"github.com/flosch/pongo2"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl := pongo2.Must(pongo2.FromFile("templates/index.html"))
	err := tpl.ExecuteWriter(pongo2.Context{"name": "Shuji!"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
