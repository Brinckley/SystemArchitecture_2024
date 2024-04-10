package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func writeJsonFromResponse(w http.ResponseWriter, status int, r *http.Response) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	_, err := io.Copy(w, r.Body)
	return err
}

func writeJson(w http.ResponseWriter, status int, content any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(content)
}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received request method: %s, header: %s, body: %s\n", r.Method, r.Header, r.Body)
		next.ServeHTTP(w, r)
	})
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			logError := fmt.Errorf("unable to write data error : %s", err)
			log.Println(logError)
			err := writeJson(w, http.StatusBadRequest, json.NewEncoder(w).Encode(logError))
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}
