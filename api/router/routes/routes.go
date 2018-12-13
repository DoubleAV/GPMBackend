package routes

import (
	"net/http"
)

//Middleware - this is the main middleware for the application
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes -
func GetRoutes() Routes {

	return Routes{
		Route{
			Name:        "HealthCheck",
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: Health(),
		},
		Route{
			Name:        "Login",
			Method:      "POST",
			Pattern:     "/auth/login",
			HandlerFunc: Login(),
		},
	}
}
