package routes

import (
	"github.com/gorilla/mux"
	"github.com/md/go-auth/internal/controllers"
)

func Router(r *mux.Router) {
	// r.HandleFunc("/", controllers.Home)

	r.HandleFunc("/google_login", controllers.GoogleLogin)
	r.HandleFunc("/google_callback", controllers.GoogleCallback)
}
