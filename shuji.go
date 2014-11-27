package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/vimrus/shuji/controllers"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", controllers.Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
