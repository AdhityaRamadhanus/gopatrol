package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/handlers"
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"github.com/AdhityaRamadhanus/gopatrol/mongo"
	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
)

var (
	app         *api.Api
	srv         *http.Server
	accessToken string
)

func clearDatabase(session *mgo.Session) {
	checkersColl := session.DB(config.DatabaseName).C("Checkers")
	usersColl := session.DB(config.DatabaseName).C("Users")
	checkersColl.DropCollection()
	usersColl.DropCollection()
}

func TestMain(m *testing.M) {
	log.SetLevel(log.WarnLevel)
	config.SetTestingConfig()
	app = api.NewApi()
	session, _ := mgo.Dial(config.MongoURI)
	defer session.Close()

	clearDatabase(session)

	// Creating Service
	usersService := mongo.NewUsersService(session, "Users")
	checkersService := mongo.NewCheckersService(session, "Checkers")

	if err := usersService.Register("Adhitya Ramadhanus", "adhitya.ramadhanus@gmail.com", "admin", "123321"); err != nil {
		log.Println(err)
	}

	checkersHandler := &handlers.CheckersHandler{
		CheckerService: checkersService,
	}
	usersHandler := &handlers.UsersHandler{
		UsersService: usersService,
	}

	app.Handlers = append(app.Handlers, checkersHandler)
	app.Handlers = append(app.Handlers, usersHandler)
	isUnixDomainSocket := false
	app.InitHandler(isUnixDomainSocket)
	srv = app.CreateServer(isUnixDomainSocket)
	code := m.Run()
	os.Exit(code)
}

func TestLoginUsers(t *testing.T) {
	userReq := struct {
		Email    string `json:"email"`
		Password string `json:"password`
	}{
		Email:    "adhitya.ramadhanus@gmail.com",
		Password: "123321",
	}
	jsonReq, _ := json.Marshal(userReq)
	req, _ := http.NewRequest("POST", "/api/v1/users/login", bytes.NewBuffer(jsonReq))
	response := httptest.NewRecorder()
	srv.Handler.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Expected to return 200 in /users/login")
	decodedResp := map[string]interface{}{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		t.Error("Failed to decode response in /users/login")
	}
	token, ok := decodedResp["accessToken"].(string)
	if !ok {
		t.Error("Faield to get access token in /users/login")
	}
	accessToken = token
}

func TestCreateCheckers(t *testing.T) {
	checkerReq := gopatrol.TCPChecker{
		Type: "tcp",
		Name: "Testing TCP Checker",
		URL:  "localhost:6379",
	}
	jsonReq, _ := json.Marshal(checkerReq)
	req, _ := http.NewRequest("POST", "/api/v1/checkers/create", bytes.NewBuffer(jsonReq))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	response := httptest.NewRecorder()
	srv.Handler.ServeHTTP(response, req)
	log.Println(response.Body.String())

	assert.Equal(t, http.StatusCreated, response.Code, "Expected to return 200 in /users/login")
	decodedResp := map[string]interface{}{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&decodedResp); err != nil {
		t.Error("Failed to decode response in /checkers/create")
	}
	message, ok := decodedResp["message"].(string)
	if !ok {
		t.Error("Faield to get data response in /checkers/login")
	}
	assert.Equal(t, "Endpoint Created", message)
}

func TestGetCheckers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/checkers/all", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	w := httptest.NewRecorder()
	srv.Handler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected to return 200 in /checkers/all")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	srv.Handler.ServeHTTP(rr, req)
	log.Println(rr.Body.String())
	return rr
}
