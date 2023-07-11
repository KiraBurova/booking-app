package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"timezone-converter/db"
)

func TestRegister(t *testing.T) {
	// read input from payload
	// add default timeslots
	// insert into db

	t.Run("register user", func(t *testing.T) {

		db.InitDb()

		user := &User{
			Username: "username",
			Password: "password",
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(user)

		request := httptest.NewRequest(http.MethodPost, "/register", payloadBuf)
		response := httptest.NewRecorder()

		Register(response, request)

		// TODO: continue test
	})
}
