package modules

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const jsonContentType = "application/json; charset=utf-8"

func toJSON(w http.ResponseWriter, statusCode int, obj interface{}) error {
	w.Header().Add("Content-Type", jsonContentType)
	w.Header().Add("Http-Code", strconv.Itoa(statusCode))
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(obj)
}
