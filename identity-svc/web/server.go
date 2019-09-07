package web

import (
	"log"
	"net/http"
	"root/cfg"
)

func Server(app *cfg.App) {
	log.Fatal(http.ListenAndServe(":8080", Router(app)))
}
