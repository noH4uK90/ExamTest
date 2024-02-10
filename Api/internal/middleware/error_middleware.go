package middleware

import (
	"errors"
	"net/http"
)

var errorCodes = map[ErrorCode]int{
	badRequest:     http.StatusBadRequest,
	unauthorized:   http.StatusUnauthorized,
	notFound:       http.StatusNotFound,
	internalServer: http.StatusInternalServerError,
	isExist:        http.StatusConflict,
}

type appHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorMiddleware(h appHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var appErr *AppError
		err := h(writer, request)
		if err != nil {
			if !errors.As(err, &appErr) {
				appErr = InternalServer
			}

			http.Error(writer, appErr.Error(), errorCodes[appErr.Code])
		}
	}
}
