package router

import (
	"net/http"
	"sagara_backend_test/lib/custerr"
	"sagara_backend_test/lib/response"
)

var (
	errUnauthorized = &custerr.ErrChain{
		Message: "unauthorized",
		Code:    http.StatusUnauthorized,
		Type:    response.ErrUnauthorized,
	}
)
