package ecoee

import (
	"ecoee/pkg/domain/model"
	"ecoee/pkg/infrastructure/dispose"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Dispose struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Count int    `json:"count" validate:"required"`
}

func (d *Dispose) Bind(r *http.Request) error {
	validator := validator.New()
	if err := validator.Struct(d); err != nil {
		return err
	}
	return nil
}

type Response struct {
	Message string `json:"message"`
}

type Registry struct {
	disposeRepository *dispose.Repository
}

func NewRegistry(disposeRepository *dispose.Repository) *Registry {
	return &Registry{
		disposeRepository: disposeRepository,
	}
}

func (r *Registry) Register(router *chi.Mux) {
	router.Group(func(router chi.Router) {
		router.Get("/", r.base)
	})
	router.Group(func(router chi.Router) {
		router.Post("/dispose", r.dispose)
	})
}

func (r *Registry) base(w http.ResponseWriter, req *http.Request) {
	render.JSON(w, req, Response{Message: "ecoee is working fine 👍"})
}

func (r *Registry) dispose(w http.ResponseWriter, req *http.Request) {
	dispose := &Dispose{}
	if err := render.Bind(req, dispose); err != nil {
		slog.Error(fmt.Sprintf("failed to bind request: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newDispose := model.Dispose{
		ID:    dispose.ID,
		Name:  dispose.Name,
		Count: dispose.Count,
	}
	result, err := r.disposeRepository.Save(req.Context(), newDispose)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to save dispose: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := Dispose{ID: result.ID, Name: result.Name, Count: result.Count}
	render.JSON(w, req, res)
}
