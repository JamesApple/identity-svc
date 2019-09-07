package web

import (
	"net/http"
	"root/cfg"
	"root/web/handlers"
)

func Routes() []Route {
	return []Route{
		{
			Verb:    "POST",
			Path:    "/",
			Handler: handlers.PostAccount,
		},
	}
}

type Route struct {
	Verb    string
	Path    string
	Handler func(*cfg.App) http.HandlerFunc
}
