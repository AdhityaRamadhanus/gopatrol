package render

import (
	"encoding/json"
	"net/http"
)

type JSON map[string]interface{}

func WriteJSON(res http.ResponseWriter, code int, v interface{}) error {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(code)
	return json.NewEncoder(res).Encode(v)
}
