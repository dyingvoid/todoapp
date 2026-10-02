package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	core_http_request "github.com/dyingvoid/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/dyingvoid/todoapp/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks godoc
//
// @Summary      List tasks
// @Description  Returns a list of tasks, optionally filtered by author user ID
// @Tags         tasks
// @Produce      json
// @Param        user_id query int false "Filter by author user ID"
// @Param        limit   query int false "Maximum number of tasks to return"
// @Param        offset  query int false "Number of tasks to skip"
// @Success      200 {array} TaskDTOResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /tasks [get]
func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID/limit/offset query params",
		)
		return
	}

	tasks, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks",
		)
		return
	}

	dtos := GetTasksResponse(tasksFromDomains(tasks))
	responseHandler.JSONResponse(dtos, http.StatusOK)
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*int64, *int, *int, error) {
	const (
		userIDQueryParam = "user_id"
		limitQueryParam  = "limit"
		offsetQueryParam = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam[int64](r, userIDQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}
	limit, err := core_http_request.GetIntQueryParam[int](r, limitQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParam[int](r, offsetQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return userID, limit, offset, nil
}
