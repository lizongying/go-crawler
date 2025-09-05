package pkg

import (
	"net/http"
)

// Route is a http.Handler that knows the mux pattern
// under which it will be registered.
type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

type MockServer interface {
	// Run starts the server.
	// It returns an error if the server fails to start or encounters a fatal issue.
	Run() error

	// AddRoutes registers one or more custom routes to the server.
	AddRoutes(...Route)

	// AddDefaultRoutes registers the built-in default routes.
	AddDefaultRoutes()

	// GetRoutes returns a list of route names currently registered in the server.
	GetRoutes() []string
}

type NewRoute func(logger Logger) Route
