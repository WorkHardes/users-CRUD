package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/users-CRUD/internal/domain"
	"github.com/users-CRUD/pkg/logger"
)

func (h *Handler) initUserRouters(api *mux.Router) {
	r := api.PathPrefix("/users").Subrouter()
	r.HandleFunc("/detail/{id}", h.getUser).Methods("GET")
	r.HandleFunc("/create", h.createUser).Methods("POST")
	r.HandleFunc("/detail/{id}", h.updateUser).Methods("PUT")
	r.HandleFunc("/detail/{id}", h.deleteUser).Methods("DELETE")
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user domain.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	ctx := context.Background()
	err := h.services.Users.Create(ctx, user)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(domain.Detail{Detail: "Failed create user!"})
		return
	}

	w.WriteHeader(201)
	response := domain.HTTP{Status: "Success!"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	ctx := context.Background()
	user, err := h.services.Users.GetById(ctx, id)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(domain.Detail{Detail: "User not found!"})
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user domain.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = mux.Vars(r)["id"]

	ctx := context.Background()
	err := h.services.Users.Update(ctx, user)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(domain.Detail{Detail: "Failed update user!"})
		return
	}

	response := domain.HTTP{Status: "Success!"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	ctx := context.Background()
	err := h.services.Users.Delete(ctx, id)
	if err != nil {
		logger.Error(err)
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(domain.Detail{Detail: "User not found!"})
		return
	}

	response := domain.HTTP{Status: "Success!"}
	json.NewEncoder(w).Encode(response)
}
