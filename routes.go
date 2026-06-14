package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NingChat is UP"))
	})
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
		http.Redirect(w, r, "https://inkdrop.quanthai.net/register", http.StatusSeeOther)
	})
}
