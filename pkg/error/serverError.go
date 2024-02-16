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

func HttpResponseError(w http.ResponseWriter, responseError ResponseError) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(responseError.GetCode()))

	resp := make(map[string]string)

	resp["message"] = responseError.Error()
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
