package routes

import (
	"api/src/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequerAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRouters
	routes = append(routes, loginRouter)

	for _, route := range routes {
		if route.RequerAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			fmt.Println("caiu aqui")
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
