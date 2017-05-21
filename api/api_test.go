package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	"gopkg.in/mgo.v2/bson"
)

var (
	app         *api.Api
	srv         *http.Server
	accessToken string
)

func clearDatabase(session *mgo.Session) {
	checkersColl := session.DB(config.DatabaseName).C("Checkers")
	usersColl := session.DB(config.DatabaseName).C("Users")
	checkersColl.RemoveAll(bson.M{})
	usersColl.RemoveAll(bson.M{})
}

func TestMain(m *testing.M) {
	log.SetLevel(log.WarnLevel)
	config.SetTestingConfig()
	app = api.NewApi()
	session, err := mgo.Dial(config.MongoURI)
	if err != nil {
		log.Error("Failed to Connect to MongoDB")
		os.Exit(1)
	}
	defer session.Close()

	clearDatabase(session)

	// Creating Service
	usersService := mongo.NewUsersService(session, "Users")
	checkersService := mongo.NewCheckersService(session, "Checkers")
	eventsService := mongo.NewEventService(session, "Events")

	if err := usersService.Register("Adhitya Ramadhanus", "adhitya.ramadhanus@gmail.com", "admin", "123321"); err != nil {
		log.Println(err)
	}

	checkersHandler := &handlers.CheckersHandler{
		CheckerService: checkersService,
	}
	usersHandler := &handlers.UsersHandler{
		UsersService: usersService,
	}
	eventsHandler := &handlers.EventsHandlers{
		EventService: eventsService,
	}

	app.Handlers = append(app.Handlers, checkersHandler)
	app.Handlers = append(app.Handlers, usersHandler)
	app.Handlers = append(app.Handlers, eventsHandler)
	isUnixDomainSocket := false
	app.InitHandler(isUnixDomainSocket)
	srv = app.CreateServer(isUnixDomainSocket)
	code := m.Run()
	os.Exit(code)
}

func TestLoginUsers(t *testing.T) {
	userReq := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

func TestCheckers(t *testing.T) {
	testCases := []struct {
		RequestBody    interface{}
		Method         string
		Path           string
		ExpectedStatus int
	}{
		{
			RequestBody: gopatrol.TCPChecker{
				Type: "tcp",
				Name: "Testing TCP Checker",
				URL:  "localhost:6379",
			},
			Method:         "POST",
			Path:           "/api/v1/checkers/create",
			ExpectedStatus: 201,
		},
		{
			Method:         "GET",
			Path:           "/api/v1/checkers/all",
			ExpectedStatus: 200,
		},
		{
			Method:         "GET",
			Path:           "/api/v1/checkers/testing-tcp-checker",
			ExpectedStatus: 200,
		},
		{
			RequestBody:    nil,
			Method:         "POST",
			Path:           "/api/v1/checkers/testing-tcp-checker/delete",
			ExpectedStatus: 200,
		},
	}

	for _, test := range testCases {
		t.Logf("Testing %s %s", test.Method, test.Path)
		var httpReq *http.Request
		switch test.Method {
		case "POST":
			jsonReq, _ := json.Marshal(test.RequestBody)
			req, _ := http.NewRequest(test.Method, test.Path, bytes.NewBuffer(jsonReq))
			httpReq = req
		case "GET":
			req, _ := http.NewRequest(test.Method, test.Path, nil)
			httpReq = req
		}

		httpReq.Header.Set("Authorization", "Bearer "+accessToken)
		response := httptest.NewRecorder()
		srv.Handler.ServeHTTP(response, httpReq)
		log.Println(response.Body.String())
		assert.Equal(t, test.ExpectedStatus, response.Code, fmt.Sprintf("Expected to return %d in %s", test.ExpectedStatus, test.Path))
		assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"), "Expected JSON")
		decoder := json.NewDecoder(response.Body)
		decodedResp := map[string]interface{}{}
		err := decoder.Decode(&decodedResp)
		assert.NoError(t, err, "Expected No Error in parsing JSON")
	}
}

func TestEvents(t *testing.T) {
	testCases := []struct {
		RequestBody    interface{}
		Method         string
		Path           string
		ExpectedStatus int
	}{
		{
			Method:         "GET",
			Path:           "/api/v1/events/all",
			ExpectedStatus: 200,
		},
	}

	for _, test := range testCases {
		t.Logf("Testing %s %s", test.Method, test.Path)
		var httpReq *http.Request
		switch test.Method {
		case "POST":
			jsonReq, _ := json.Marshal(test.RequestBody)
			req, _ := http.NewRequest(test.Method, test.Path, bytes.NewBuffer(jsonReq))
			httpReq = req
		case "GET":
			req, _ := http.NewRequest(test.Method, test.Path, nil)
			httpReq = req
		}

		httpReq.Header.Set("Authorization", "Bearer "+accessToken)
		response := httptest.NewRecorder()
		srv.Handler.ServeHTTP(response, httpReq)
		log.Println(response.Body.String())
		assert.Equal(t, test.ExpectedStatus, response.Code, fmt.Sprintf("Expected to return %d in %s", test.ExpectedStatus, test.Path))
		assert.Equal(t, "application/json; charset=utf-8", response.Header().Get("Content-Type"), "Expected JSON")
		decoder := json.NewDecoder(response.Body)
		decodedResp := map[string]interface{}{}
		err := decoder.Decode(&decodedResp)
		assert.NoError(t, err, "Expected No Error in parsing JSON")
	}
}
