package routes

import (
	"autenticacao-jwt/logger"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		if route.Protected {
			protectedRouter := mux.NewRouter()
			protectedRouter.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)

			router.PathPrefix(route.Pattern).Handler(negroni.New(
				negroni.HandlerFunc(validateTokenMiddleware),
				negroni.Wrap(protectedRouter),
			))

		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}

	}

	log.Println("Now listening...")

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		true,
		"/index",
		IndexHandler,
	},
	Route{
		"Login",
		"POST",
		false,
		"/login",
		LoginHandler,
	},
}

func validateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

}
