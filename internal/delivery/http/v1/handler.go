package v1

import (
	"github.com/gorilla/mux"
	"github.com/users-CRUD/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *mux.Router) {
	v1 := api.PathPrefix("/v1").Subrouter()
	h.initUserRouters(v1)
}
