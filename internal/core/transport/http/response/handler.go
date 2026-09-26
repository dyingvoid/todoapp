package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_error "github.com/dyingvoid/todoapp/internal/core/errors"
	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
	core_logger "github.com/dyingvoid/todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(
	log *core_logger.Logger,
	w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponseHandler) JSONResponse(
	responseBody any,
	statusCode int,
) {
	h.w.WriteHeader(statusCode)
	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFn      func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFn = h.log.Warn
	case errors.Is(err, core_error.ErrNotFound):
		statusCode = http.StatusNotFound
		logFn = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFn = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFn = h.log.Error
	}

	logFn(msg, zap.Error(err))

	h.errorResponse(statusCode, err, msg, logFn)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)
	h.errorResponse(statusCode, err, msg, h.log.Error)
}

func (h *HTTPResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
	logFn func(string, ...zap.Field),
) {
	logFn(msg, zap.Error(err))
	h.w.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(
		response,
		statusCode,
	)
}
