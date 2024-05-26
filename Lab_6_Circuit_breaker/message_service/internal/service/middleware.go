package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, content any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(content)
}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("message service received request method: %s, header: %s, body: %s\n", r.Method, r.Header, r.Body)
		next.ServeHTTP(w, r)
	})
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(fmt.Errorf("some error happened in msg service : %s", err))
			err := writeJson(w, http.StatusBadRequest, err)
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}
