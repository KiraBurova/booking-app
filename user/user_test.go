package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"timezone-converter/db"
)

func TestRegister(t *testing.T) {
	t.Run("register user", func(t *testing.T) {
		db.InitDb()
		db := ConnectDB{db: db.DbInstance}

		user := &User{
			Username: "user_created_from_test",
			Password: "user_created_from_test",
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(user)

		request := httptest.NewRequest(http.MethodPost, "/register", payloadBuf)
		response := httptest.NewRecorder()

		Register(response, request)

		row := db.queryRow("SELECT * FROM users WHERE username=?", "user_created_from_test")

		// for now
		fmt.Println(row)
	})
}
