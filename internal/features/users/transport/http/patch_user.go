package users_transport_http

import (
	"net/http"

	"github.com/dyingvoid/todoapp/internal/core/domain"
	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	core_http_request "github.com/dyingvoid/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/dyingvoid/todoapp/internal/core/transport/http/response"
	core_http_types "github.com/dyingvoid/todoapp/internal/core/transport/http/types"
	core_http_utils "github.com/dyingvoid/todoapp/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get id path value",
		)
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
	}
	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(r PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName:    r.FullName.ToDomain(),
		PhoneNumber: r.PhoneNumber.ToDomain(),
	}
}
