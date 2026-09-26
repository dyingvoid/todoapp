package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	log := core_logger.FromContext(r.Context())
	log.Debug("invoke CreateUserHandler")

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error("decode request", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
