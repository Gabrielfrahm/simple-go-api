package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRouter = Route{
	Uri:                  "/login",
	Method:               http.MethodPost,
	Function:             controllers.Login,
	RequerAuthentication: false,
}
