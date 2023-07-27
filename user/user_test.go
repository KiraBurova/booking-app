package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"timezone-converter/db"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("register user", func(t *testing.T) {
		db.InitDb("testUsers.db")
		repo := NewRepository(db.DbInstance)

		user := &User{
			Username: "user_created_from_test",
			Password: "user_created_from_test",
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(user)

		request := httptest.NewRequest(http.MethodPost, "/register", payloadBuf)
		response := httptest.NewRecorder()

		Register(response, request)

		row, err := repo.GetByUsername("user_created_from_test")

		if err != nil {
			t.Fail()
		}

		assert.Equal(t, row.Username, "user_created_from_test", "User with username of user_created_from_test should be returned")
	})
}
