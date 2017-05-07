package helper

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"strings"
)

func WriteGzip(res http.ResponseWriter, req *http.Request, code int, v interface{}) error {
	res.Header().Set("Content-Type", "application/json")
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		res.Header().Set("Content-Encoding", "gzip")
		gWriter := gzip.NewWriter(res)
		defer gWriter.Close()
		res.WriteHeader(code)
		return json.NewEncoder(gWriter).Encode(v)
	}
	return WriteJSON(res, code, v)
}

func WriteGzipBytes(res http.ResponseWriter, req *http.Request, code int, v []byte) error {
	res.Header().Set("Content-Type", "application/json")
	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		res.Header().Set("Content-Encoding", "gzip")
		gWriter := gzip.NewWriter(res)
		defer gWriter.Close()
		res.WriteHeader(code)
		_, err := gWriter.Write(v)
		return err
	}
	return WriteJSON(res, code, v)
}
