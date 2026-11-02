package core_http_middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func RequestID() Middleware {
	const requestIDKey = "x-request-id"

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDKey)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDKey, requestID)
			w.Header().Set(requestIDKey, requestID)

			next.ServeHTTP(w, r)
		})
	}
}
