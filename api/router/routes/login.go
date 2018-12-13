package routes

import (
	"log"
	"net/http"
)

// Login - login route
func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		log.Println(email)
		log.Println(password)
	}
}
