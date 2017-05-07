package helper

import (
	"encoding/json"
	"net/http"
	"time"
)

func WriteJSON(res http.ResponseWriter, code int, v interface{}) error {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	var response map[string]interface{}
	switch code {
	case http.StatusOK:
		response = map[string]interface{}{
			"timestamp": time.Now().UTC().String(),
			"data":      v,
		}
	case http.StatusInternalServerError:
		response = map[string]interface{}{
			"timestamp": time.Now().UTC().String(),
			"error":     v,
		}

	default:
		response = map[string]interface{}{
			"timestamp": time.Now().UTC().String(),
			"message":   v,
		}
	}
	return json.NewEncoder(res).Encode(response)

}
