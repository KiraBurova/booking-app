package auth

import (
	"context"
	"net/http"
	"timezone-converter/db"
)

// ctxKeyRequestID defines the request ID key to be stored/retrieved from the context.
type ctxKeySessionId struct{}

func WithSessionId(ctx context.Context, cookie string) context.Context {
	return context.WithValue(ctx, ctxKeySessionId{}, cookie)
}

func SessionId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := getCookie(r)

		if err != nil {
			repo := NewRepository(db.DbInstance)

			cookieExists, _ := repo.Get(cookie.Value)

			if cookieExists == true {
				ctx := WithSessionId(r.Context(), cookie.Value)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
	})
}
