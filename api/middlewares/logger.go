package middlewares

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

func HTTPReqLogger(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		nextHandler.ServeHTTP(res, req)
		log.WithFields(log.Fields{
			"method":    req.Method,
			"resp time": time.Since(start),
		}).Info(req.RequestURI)
	})
}
