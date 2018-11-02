package service

import (
	"net/http"
	"strings"
)

// Authenticate wraps HTTP requests with Authentication
func Authenticate(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokens := strings.Split(name, "_")
		if tokens[0] == "" {
			inner.ServeHTTP(w, r)
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}

func auth() {

}
