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
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		repo := NewRepository(db.DbInstance)

		cookieExists, cookieExistsError := repo.Get(cookie.Value)

		if cookieExistsError != nil || cookieExists == false {
			http.Error(w, "Authentication error.", http.StatusUnauthorized)
			return
		}

		ctx := WithSessionId(r.Context(), cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
