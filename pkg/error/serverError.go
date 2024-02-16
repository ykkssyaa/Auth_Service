package server_error

import (
	"encoding/json"
	"net/http"
)

type Status int

type ResponseError struct {
	Code    Status
	Message string
}

func (re ResponseError) Error() string {
	return re.Message
}

func (re ResponseError) GetCode() Status {
	return re.Code
}

func HttpResponseError(w http.ResponseWriter, message string, httpStatusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := make(map[string]string)

	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
