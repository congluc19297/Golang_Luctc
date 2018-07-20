package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// Injects a request ID into the context of each request
	r.Use(middleware.RequestID)
	// Logs the start and end of each request with the elapsed processing time
	r.Use(middleware.Logger)

	r.Get("/pets", funcIndex)

	http.ListenAndServe(":3333", r)
}

// swagger:route GET /pets pets users listPets
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//       api_key:
//       oauth: read, write
//     Responses:
//       default: genericError
//       200: someResponse
//       422: validationError
func funcIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
