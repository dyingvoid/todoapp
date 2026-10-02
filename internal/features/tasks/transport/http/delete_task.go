package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	core_http_request "github.com/dyingvoid/todoapp/internal/core/transport/http/request"
	core_http_response "github.com/dyingvoid/todoapp/internal/core/transport/http/response"
)

// DeleteTask godoc
//
// @Summary      Delete task
// @Description  Deletes a task by its ID
// @Tags         tasks
// @Param        id path int true "Task ID"
// @Success      204 "No Content"
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /tasks/{id} [delete]
func (h *TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task 'id' path value",
		)
		return
	}

	if err := h.tasksService.DeleteTask(ctx, taskID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete task",
		)
		return
	}

	responseHandler.NoContentResponse()
}
