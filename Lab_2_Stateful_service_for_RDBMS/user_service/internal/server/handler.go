package server

import (
	"fmt"
	"net/http"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJson(w, http.StatusBadRequest, err)
		}
	}
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.getAccount(w, r)
	case http.MethodPost:
		return s.createAccount(w, r)
	case http.MethodDelete:
		return s.deleteAccount(w, r)
	case http.MethodPut:
		return nil
	}
	return fmt.Errorf("[ERR] invalid method : %s", r.Method)
}

func (s *ApiServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
