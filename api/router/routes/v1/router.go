package v1

import (
	"net/http"

	Routes "github.com/backend/api/router/routes"
)

// Middleware - this is the middleware for v1 routes
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1 routes
func GetRoutes() (SubRoute map[string]Routes.SubRoutePackage) {
	SubRoute = map[string]Routes.SubRoutePackage{
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name:        "V1HealthRotue",
					Method:      "GET",
					Pattern:     "/health",
					HandlerFunc: Health(),
				},
			},
			Middleware: Middleware,
		},
	}

	return SubRoute
}
