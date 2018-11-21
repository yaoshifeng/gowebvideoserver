package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:login_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	log.Fatal(http.ListenAndServe(":8080", r))
}
