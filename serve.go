package main

import (
	"go-jwt/handler"
	"go-jwt/middleware"
	"net/http"
)

func main() {

	loginHandler := http.HandlerFunc(handler.LoginHandler)
	http.Handle("/login", loginHandler)
	http.Handle("/home", middleware.AuthorizeJwt(handler.HomeHandler))

	_ = http.ListenAndServe(":2345", nil)
}
