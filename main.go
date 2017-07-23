package main

//  Generate RSA signing files via shell (adjust as needed):
//
//  $ openssl genrsa -out app.rsa 1024
//  $ openssl rsa -in app.rsa -pubout > app.rsa.pub
//
// Code borrowed and modified from the following sources:
// https://www.youtube.com/watch?v=dgJFeqeXVKw
// https://goo.gl/ofVjK4
// https://github.com/dgrijalva/jwt-go
// http://www.giantflyingsaucer.com/blog/?p=5994

import (
	"autenticacao-jwt/routes"
	"log"
	"net/http"
)

//https://thenewstack.io/make-a-restful-json-api-go/
func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
