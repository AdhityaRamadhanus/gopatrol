package middlewares

import (
	"errors"
	"net/http"

	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	"github.com/AdhityaRamadhanus/gopatrol/api/render"
	"github.com/AdhityaRamadhanus/gopatrol/config"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	scopeMap = map[string]int{
		"admin":  0,
		"member": 1,
	}
)

func AuthenticateToken(nextHandler http.HandlerFunc, scope int) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader, ok := req.Header["Authorization"]
		if !ok || len(authHeader) == 0 {
			render.WriteJSON(res, http.StatusUnauthorized, render.JSON{
				"error": "Authorization Header not Present",
			})
			return
		}
		cred, err := helper.ParseAuthorizationHeader(authHeader[0], "Bearer")
		if err != nil {
			render.WriteJSON(res, http.StatusUnauthorized, render.JSON{
				"error": err,
			})
			return
		}
		token, err := jwt.Parse(cred, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method")
			}
			return []byte(config.JwtSecret), nil
		})
		if err != nil {
			render.WriteJSON(res, http.StatusUnauthorized, render.JSON{
				"error": err,
			})
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid && scopeMap[claims["role"].(string)] < scope {
			nextHandler(res, req)
		} else {
			render.WriteJSON(res, http.StatusUnauthorized, render.JSON{
				"error": "Cannot authorize token",
			})
			return
		}
	})
}
