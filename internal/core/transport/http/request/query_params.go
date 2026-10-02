package core_http_request

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/dyingvoid/todoapp/internal/core/errors"
)

func GetIntQueryParam[T int | int64](r *http.Request, key string) (*T, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid integer: %v: %w",
			param,
			key,
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	result := T(val)
	return &result, nil
}
