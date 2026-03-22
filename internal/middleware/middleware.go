package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

type contextKey string

const (
	// UserIDKey is the key to store the user ID in the context
	UserIDKey contextKey = "userID"
)

// Logger logs the incoming request method, path and its duration
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("[Logger] %s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Recovery catches panics and prevents the server from crashing
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Recovery] Panic caught: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Auth simple middleware acting as an API key verifier
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey != "secret-key" {
			log.Printf("[Auth] Unauthorized access attempt")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Inject context with mock user ID
		ctx := context.WithValue(r.Context(), UserIDKey, "user-1234")
		reqWithCtx := r.WithContext(ctx)

		next.ServeHTTP(w, reqWithCtx)
	})
}
