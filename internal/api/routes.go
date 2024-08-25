package api

import (
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

// SetupRoutes configures the routes for our API
func SetupRoutes() {

	// Create a new pocketbase instance
	app := pocketbase.New()

	routes := []Route{
		{Path: "/", Method: "GET", Description: "List all available routes"},
		{Path: "/v1/about/", Method: "GET", Description: "Get information about the API"},
		{Path: "/v1/new/:len", Method: "GET", Description: "Returns a list with numbers of length"},
	}

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/", RootHandler(routes))
		e.Router.GET("/v1/about/", AboutHandler)
		e.Router.GET("/v1/new/:len/", GenerateNumbersHandler(app))

		return nil
	})

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
