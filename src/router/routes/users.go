package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Route{
	{
		Uri:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequerAuthentication: true,
	},
	{
		Uri:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.ListUser,
		RequerAuthentication: false,
	},
	{
		Uri:                  "/users/{userId}",
		Method:               http.MethodGet,
		Function:             controllers.ListOneUser,
		RequerAuthentication: false,
	},
	{
		Uri:                  "/users/{userId}",
		Method:               http.MethodPut,
		Function:             controllers.UpdateUser,
		RequerAuthentication: false,
	},
	{
		Uri:                  "/users/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequerAuthentication: false,
	},
}
