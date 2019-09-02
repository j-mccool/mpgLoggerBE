package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"../controller"
)

// Router creates and returns a new router
func Router(ctlr controller.Controller) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")
	r.HandleFunc("/health", controller.HealthCheck).Methods("GET")

	return r
}

// StartWithCORS starts router with CORS in case we're running this local with a local front end
func StartWithCORS(r *mux.Router) {
	headers := handlers.AllowHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowMethods([]string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fata(http.ListenAndServe(":8080", hanlders.CORS(headers, methods, origins)(r)))
}
