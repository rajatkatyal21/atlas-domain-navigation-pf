package models

import (
	"encoding/json"
	"net/http"
)

// ResponseWriter is the interface for Writing to the Response.
type RespWriter interface {
	SendHttpResponse(w http.ResponseWriter)
}

// Status enum
type Status string

const (
	OK  Status = "OK"
	NOK Status = "NOK"
)

// Response is the format in which the response will be sent.
type Response struct {
	response interface{}
	code     int
}

//Error Response Struct
type ErrorResponse struct {
	Reason string
	Status Status
	Code   int
}

// NewError is used for creating the ErrorResponse.
func NewError(reason string, status Status, code int) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Reason: reason,
		Code:   code,
	}
}

// NewResponseSetter is a constructor to create the RespWriter.
// It takes res and response code as paramete
// Returns the RespWriter.
func NewResponseSetter(res interface{}, code int) RespWriter {
	return &Response{
		response: res,
		code:     code,
	}
}

// SendHttpResponse is a wrapper that can be used as common function to send back the http response.
func (r *Response) SendHttpResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(r.response)
	if err != nil {
		//In case there is error while Marshalling.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(r.code)
	w.Write(res)

}
