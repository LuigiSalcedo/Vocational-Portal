package core

import (
	"fmt"
	"net/http"
)

// Request response struct
type HttpResponse struct {
	Code int `json:"code"`
}

// Request error response struct
type HttpError struct {
	HttpResponse
	Error string `json:"error"`
}

// Request special respose structu
type HttpSpecial struct {
	HttpResponse
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Communs responses
var (
	OK         = HttpResponse{http.StatusOK}                  // 200
	Created    = HttpResponse{http.StatusCreated}             // 201
	BadRequest = HttpResponse{http.StatusBadRequest}          // 400
	NotFound   = HttpResponse{http.StatusNotFound}            // 404
	Internal   = HttpResponse{http.StatusInternalServerError} // 500
)

// Communs errors
var (
	PathError = &HttpError{
		HttpResponse: BadRequest,
		Error:        "the path is malformed",
	}

	NotFoundError = &HttpError{
		HttpResponse: NotFound,
		Error:        "resource not found",
	}

	JsonError = &HttpError{
		HttpResponse: BadRequest,
		Error:        "json is not valid",
	}

	QueryError = &HttpError{
		HttpResponse: BadRequest,
		Error:        "query param is malformed",
	}

	InternalError = &HttpError{
		HttpResponse: Internal,
		Error:        "internal server error",
	}
)

// function for create query errors using a specific query param
func NewQueryError(param string) *HttpError {
	return NewError(http.StatusBadRequest, fmt.Sprintf("query param named '%s' is malformed", param))
}

// function for create http error
func NewError(code int, msg string) *HttpError {
	return &HttpError{
		HttpResponse: HttpResponse{Code: code},
		Error:        msg,
	}
}
