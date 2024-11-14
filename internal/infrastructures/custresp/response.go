package custresp

import (
	"database/sql"
	"errors"
	"net/http"
	"sagara_backend_test/internal/interfaces/dao"
	"sagara_backend_test/lib/custerr"
	"sagara_backend_test/lib/response/rest"
	"sagara_backend_test/pkg/constants"
	"sagara_backend_test/pkg/constants/errorcode"
)

func CustomErrorResponse(err error) (*rest.JSONResponse, error) {
	resp := rest.NewJSONResponse()
	if err == nil {
		return resp, nil
	}

	var e *custerr.ErrChain
	switch {
	case errors.As(err, &e):
		errCause := e.Cause
		if errCause != nil {
			err = errCause
		}

		message := e.Message
		if message == constants.EmptyString {
			message = err.Error()
		}

		resp.SetCode(getErrorCode(e))
		resp.Error = &rest.ErrorResponse{
			ErrorCode:    e.Code,
			ErrorMessage: message,
		}

		return resp, nil
	case errors.Is(err, dao.ErrNoResult), errors.Is(err, sql.ErrNoRows):
		resp.SetCode(http.StatusBadRequest)
		resp.Error = &rest.ErrorResponse{
			ErrorCode:    errorcode.NotFound.Code,
			ErrorMessage: errorcode.NotFound.Message,
		}
		return resp, nil
	default:
		return resp.SetError(err).SetMessage(err.Error()), nil
	}
}

func getErrorCode(err *custerr.ErrChain) int {
	switch err.Type {
	case ErrTooManyRequest:
		return http.StatusTooManyRequests
	case ErrRequestTooEarly:
		return http.StatusTooEarly
	case ErrInvalidRequest:
		return http.StatusNotAcceptable
	default:
		// if it not found, then it will use library to define the http code
		return rest.GetErrorCode(err)
	}
}
