package web

import (
	"log"
	"net/http"
	"root/cfg"

	"github.com/gorilla/mux"
)

func Router(app *cfg.App) http.Handler {
	r := mux.NewRouter()
	for _, route := range Routes() {
		r.Path(route.Path).Methods(route.Verb).HandlerFunc(route.Handler(app))
	}
	r.Use(InstrumentRoute)
	return r
}

func InstrumentRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %v", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			val := recover()
			switch err := val.(type) {
			case nil:
				return
			case error:
				log.Print(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			default:
				log.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, req)
	})
}
