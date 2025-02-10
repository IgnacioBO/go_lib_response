package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func errors(msg string, code int) Response {
	return &ErrorResponse{
		Status:  code,
		Message: msg,
	}
}

func InternalServerError(msg string) Response {
	return errors(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return errors(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return errors(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return errors(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return errors(msg, http.StatusBadRequest)
}

//Metodos para implementar los del interface

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e) //Marshall del objeto errorResponse para transformarlo en json
}

func (e *ErrorResponse) GetData() interface{} {
	return nil
}
