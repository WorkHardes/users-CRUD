package delivery

import (
	"github.com/gorilla/mux"
	"github.com/users-CRUD/internal/config"
	v1 "github.com/users-CRUD/internal/delivery/http/v1"
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

func (h *Handler) Init(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *mux.Router) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.PathPrefix("/api").Subrouter()
	handlerV1.Init(api)
}
