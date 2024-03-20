package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/db"
)

type ApiServer struct {
	Storage db.Storage
	Host    string
	Port    string
}

func NewApiServer(port string, storage db.Storage) *ApiServer {
	return &ApiServer{
		Storage: storage,
		Port:    port,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	router.HandleFunc("/account", makeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.getAccount)).Methods(http.MethodGet)
	router.HandleFunc("/account", makeHTTPHandleFunc(s.createAccount)).Methods(http.MethodPost)
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.modifyAccount)).Methods(http.MethodPut)
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.deleteAccount)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.Port, router)
	if err != nil {
		return
	}
}
