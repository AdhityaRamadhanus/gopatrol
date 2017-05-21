package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/render"
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	UsersService gopatrol.UsersService
}

func (h *UsersHandler) AddRoutes(router *mux.Router, isUnixDomain bool) {
	// if !isUnixDomain {
	router.HandleFunc("/users/register", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/login", h.Login).Methods("POST")
	// }
}

func (h *UsersHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	// Read Body, limit to 1 MB //
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToReadBody,
		})
		return
	}
	if err := req.Body.Close(); err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToCloseBody,
		})
		return
	}
	user := struct {
		Name     string `json:"name" valid:"required"`
		Email    string `json:"email" valid:"required,email"`
		Role     string `json:"role" valid:"required"`
		Password string `json:"password" valid:"required"`
	}{}

	// Deserialize
	if err := json.Unmarshal(body, &user); err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToUnmarshalJSON,
		})
		return
	}

	if len(user.Password) < 8 {
		render.WriteJSON(res, http.StatusBadRequest, render.JSON{
			"error": "Password must be 8 length or greater",
		})
		return
	}

	if ok, err := govalidator.ValidateStruct(user); !ok || err != nil {
		render.WriteJSON(res, http.StatusBadRequest, render.JSON{
			"error": api.ErrFailedToValidateStruct,
		})
		return
	}

	if err := h.UsersService.Register(user.Name, user.Email, user.Role, user.Password); err != nil {
		log.WithError(err).Error("Failed to create user")
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrInternalServerError,
		})
		return
	}
	render.WriteJSON(res, http.StatusCreated, render.JSON{
		"message": "User created",
	})
}

func (h *UsersHandler) Login(res http.ResponseWriter, req *http.Request) {
	// Read Body, limit to 1 MB //
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToReadBody,
		})
		return
	}
	if err := req.Body.Close(); err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToCloseBody,
		})
		return
	}

	user := struct {
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required"`
	}{}

	// Deserialize
	if err := json.Unmarshal(body, &user); err != nil {
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": api.ErrFailedToUnmarshalJSON,
		})
		return
	}

	if ok, err := govalidator.ValidateStruct(user); !ok || err != nil {
		render.WriteJSON(res, http.StatusBadRequest, render.JSON{
			"error": api.ErrFailedToValidateStruct,
		})
		return
	}

	dbUser, err := h.UsersService.Login(user.Email, user.Password)
	if err != nil {
		render.WriteJSON(res, http.StatusNotFound, render.JSON{
			"error": "Failed to authenticate",
		})
		return
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     dbUser.Email,
		"role":      dbUser.Role,
		"timestamp": time.Now(),
	})
	tokenString, err := jwtToken.SignedString([]byte(config.JwtSecret))
	if err != nil {
		log.WithError(err).Error("Failed to create access token")
		render.WriteJSON(res, http.StatusInternalServerError, render.JSON{
			"error": "Failed to create access token",
		})
		return
	}

	render.WriteJSON(res, http.StatusOK, render.JSON{
		"accessToken": tokenString,
		"user":        dbUser,
	})
}
