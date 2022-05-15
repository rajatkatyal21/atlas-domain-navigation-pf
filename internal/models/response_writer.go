package models

import (
	"encoding/json"
	"net/http"
)

type RespWriter interface {
	SendHttpResponse(w http.ResponseWriter)
}

// Response is the format in which the response will be sent.
type Response struct {
	response interface{}
	code     int
}

// NewResponseSetter is a constructor to create the RespWriter.
// It takes res and response code as paramete
// Returns the RespWriter.
func NewResponseSetter(res interface{}, code int) RespWriter {
	return &Response{
		response: res,
		code: code,
	}
}

// SendHttpResponse can be used a common function to send back the http response.
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

