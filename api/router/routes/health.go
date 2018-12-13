package routes

import (
	"net/http"
)

// Health - health route
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy"))
	}
}
