package server

import (
	"account_service/internal/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountApiServer struct {
	Storage     repository.Storage
	AccountPort string
}

func NewAccountApiServer(port string, storage repository.Storage) *AccountApiServer {
	return &AccountApiServer{
		Storage:     storage,
		AccountPort: port,
	}
}

func (s *AccountApiServer) Run() {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	router.HandleFunc("/", makeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)
	router.HandleFunc("/", makeHTTPHandleFunc(s.createAccount)).Methods(http.MethodPost)
	router.HandleFunc("/{id}", makeHTTPHandleFunc(s.getAccount)).Methods(http.MethodGet)
	router.HandleFunc("/{id}", makeHTTPHandleFunc(s.updateAccount)).Methods(http.MethodPut)
	router.HandleFunc("/{id}", makeHTTPHandleFunc(s.deleteAccount)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.AccountPort, router)
	if err != nil {
		return
	}
}
