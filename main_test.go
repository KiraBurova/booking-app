package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddTimezone(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/timezone", strings.NewReader("Test/Timezone"))

	addTimezone(w, r)
}
