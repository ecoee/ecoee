package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Registry struct{}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Register(router *chi.Mux) {
	router.Group(func(router chi.Router) {
		router.Get("/health", r.health)
	})
}

func (r *Registry) health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
