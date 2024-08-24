package api

import "net/http"

// Chain applies a series of middlewares to a http.Handler
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

// SetupRoutes configures the routes for our API
func SetupRoutes() http.Handler {

	routes := []Route{
		{Path: "/", Method: "GET", Description: "List all available routes"},
		{Path: "/v1/about/", Method: "GET", Description: "Get information about the API"},
		{Path: "/v1/new/{length}/", Method: "GET", Description: "Returns a list with numbers of length"},
	}

	mux := http.NewServeMux()

	mux.Handle("GET /", RootHandler(routes))
	mux.HandleFunc("GET /v1/about/", AboutHandler)
	mux.HandleFunc("GET /v1/new/{length}/", GenerateNumbersHandler)

	// Apply multiple middlewares
	return Chain(mux,
		LoggerMiddleWare,
		JSONContentTypeMiddleware,
	)

}
