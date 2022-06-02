package common

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest      errorControllerResponseCode = "bad_request"
	ErrForbiddenAccess errorControllerResponseCode = "forbidden"
)

type ControllerResponse struct {
	Code    errorControllerResponseCode `json:"code"`
	Message string                      `json:"message"`
	Data    interface{}                 `json:"data"`
}

func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		Code:    ErrBadRequest,
		Message: "Bad Request",
		Data:    nil,
	}
}
