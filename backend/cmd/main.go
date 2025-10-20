package main

import (
	"hutchison-test/common"
	"hutchison-test/infrastructure"
	"hutchison-test/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	for _, route := range routes.Routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = common.CorsMiddleware(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		router.
			Methods("OPTIONS").
			Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, ANY")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusNoContent)
			}))

	}

	db := infrastructure.InitialiseDB()
	defer db.Close()

	log.Printf("Server is listening on port 5050")
	log.Fatal(http.ListenAndServe(":5050", router))
}
