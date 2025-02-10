package response

import (
	"encoding/json"
	"net/http"

	"github.com/IgnacioBO/gomicro_meta/meta"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

func success(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return &SuccessResponse{
		Message: msg,
		Status:  code,
		Data:    data,
		Meta:    meta,
	}
}

func OK(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusOK)
}

func Created(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusCreated)
}

func Accepted(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusAccepted)
}

func NotAuthoritiveInfo(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusNoContent)
}

func ResetContent(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusResetContent)
}

func PartialContent(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return success(msg, data, meta, http.StatusPartialContent)
}

//Metodos para implementar los del interface

func (successResponse *SuccessResponse) Error() string {
	return "" //Vacio porque es success
}

func (successResponse *SuccessResponse) StatusCode() int {
	return successResponse.Status
}

func (successResponse *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(successResponse) //Marshall del objeto succesRespone para transformarlo en json
}

func (successResponse *SuccessResponse) GetData() interface{} {
	return successResponse.Data
}
