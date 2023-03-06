package app

import (
	"net/http"

	"coke/app/actions"
	"coke/app/actions/home"
	"coke/app/middleware"
	"coke/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	users := actions.UsersResource{}
	root.GET("/", home.Index)
	root.GET("/users", users.List)
	root.ServeFiles("/", http.FS(public.FS()))
}
