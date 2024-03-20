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
		RequerAuthentication: false,
	},
	{
		Uri:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.ListUser,
		RequerAuthentication: true,
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
		RequerAuthentication: true,
	},
	{
		Uri:                  "/users/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequerAuthentication: true,
	},
	{
		Uri:                  "/users/{userId}/follow",
		Method:               http.MethodPost,
		Function:             controllers.FollowUser,
		RequerAuthentication: true,
	},
	{
		Uri:                  "/users/{userId}/unfollow",
		Method:               http.MethodPost,
		Function:             controllers.UnfollowUser,
		RequerAuthentication: true,
	},
	{
		Uri:                  "/users/{userId}/followers",
		Method:               http.MethodGet,
		Function:             controllers.SearchFollowers,
		RequerAuthentication: true,
	},
	{
		Uri:                  "/users/{userId}/followings",
		Method:               http.MethodGet,
		Function:             controllers.SearchFollowings,
		RequerAuthentication: true,
	},
}
