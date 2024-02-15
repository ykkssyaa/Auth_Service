package server_error

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := make(map[string]string)

	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	_, _ = w.Write(jsonResp)
}
