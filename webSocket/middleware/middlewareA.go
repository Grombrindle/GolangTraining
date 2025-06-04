// Package middleware provides HTTP middleware functionality
package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware is a type that takes an http.Handler and returns a new http.Handler
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging creates a middleware that logs request path and processing time
func Logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Record start time
			start := time.Now()

			// Defer logging until after request is handled
			defer func() {
				log.Printf("Path: %s, Duration: %v\n", r.URL.Path, time.Since(start))
			}()

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Method creates a middleware that restricts requests to a specific HTTP method
func Method(m string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if method matches
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed),
					http.StatusMethodNotAllowed)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Chain applies multiple middleware to a http.HandlerFunc in order
func Chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// Hello is a simple handler that writes "hello world" to the response
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// main sets up and starts the HTTP server with middleware
func MiddlewareA() {
	// Create handler with middleware chain
	handler := Chain(
		http.HandlerFunc(Hello),
		Method("GET"),
		Logging(),
	)

	// Register handler and start server
	http.Handle("/handler", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
