package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	// read input from payload
	// add default timeslots
	// insert into db

	t.Run("register user", func(t *testing.T) {

		user := &User{
			Username: "username",
			Password: "password",
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(user)

		request := httptest.NewRequest(http.MethodPost, "/register", payloadBuf)
		response := httptest.NewRecorder()

		Register(response, request)

		// TODO: test fails here with runtime error: invalid memory address or nil pointer dereference
	})
}
