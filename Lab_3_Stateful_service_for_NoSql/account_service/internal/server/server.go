package server

import (
	"account_service/internal/storage"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountApiServer struct {
	Storage     storage.Storage
	AccountPort string
	Ctx         context.Context
}

func NewAccountApiServer(port string, storage storage.Storage, ctx context.Context) *AccountApiServer {
	return &AccountApiServer{
		Storage:     storage,
		AccountPort: port,
		Ctx:         ctx,
	}
}

func (s *AccountApiServer) Run() {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	//router.HandleFunc("/search", makeHTTPHandleFunc(s.getAccountsByMask)).Methods(http.MethodGet)
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)
	router.HandleFunc("/accounts", makeHTTPHandleFunc(s.createAccount)).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(s.getAccount)).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(s.updateAccount)).Methods(http.MethodPut)
	router.HandleFunc("/accounts/{id}", makeHTTPHandleFunc(s.deleteAccount)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.AccountPort, router)
	if err != nil {
		return
	}
}
