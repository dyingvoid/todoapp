package users_transport_http

import (
	"net/http"

	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	core_http_request "github.com/dyingvoid/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/dyingvoid/todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
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
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
