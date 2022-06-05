package main

import (
	"github.com/julienschmidt/httprouter"

	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/user/:id", ShowUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
