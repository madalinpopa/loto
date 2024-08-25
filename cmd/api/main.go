package main

import (
	_ "github.com/madalinpopa/loto/migrations"

	"github.com/madalinpopa/loto/internal/api"
)

func main() {

	// Setup a new Mux Server and its routes
	// mux := api.SetupRoutes()

	// server := &http.Server{
	// 	Addr:    ":8000",
	// 	Handler: mux,
	// }

	// log.Println("Server starting on :8000")
	// log.Fatal(server.ListenAndServe())

	api.SetupRoutes()
}
