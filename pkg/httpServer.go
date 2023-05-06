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

type DevServer interface {
	Run() error
	AddRoutes(...Route)
	GetHost() string
}
