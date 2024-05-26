package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("account service received request method: %s, header: %s, body: %s\n", r.Method, r.Header, r.Body)
		next.ServeHTTP(w, r)
	})
}
