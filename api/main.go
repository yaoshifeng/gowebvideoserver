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

//handle->validation(1，request:校验用户名，密码是否合法, 2，是不是注册过的合法用户)->business logic走商务逻辑->response
