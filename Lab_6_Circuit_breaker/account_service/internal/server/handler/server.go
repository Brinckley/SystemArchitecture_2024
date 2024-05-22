package handler

import (
	"account_service/internal/repository"
	"account_service/internal/server/middleware"
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
	router.Use(middleware.LoggingMiddleWare)

	router.HandleFunc("/auth", middleware.MakeHTTPHandleFunc(s.getAccountIdFromToken)).Methods(http.MethodGet)
	router.HandleFunc("/search", middleware.MakeHTTPHandleFunc(s.getAccountsByMask)).Methods(http.MethodGet)

	router.HandleFunc("/signup", middleware.MakeHTTPHandleFunc(s.signUpAccount)).Methods(http.MethodPost)
	router.HandleFunc("/signin", middleware.MakeHTTPHandleFunc(s.signInAccount)).Methods(http.MethodPost)

	router.HandleFunc("/accounts", middleware.MakeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)

	router.HandleFunc("/accounts/{account_id}", middleware.MakeHTTPHandleFunc(s.getAccountById)).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account_id}", middleware.MakeHTTPHandleFunc(s.updateAccountById)).Methods(http.MethodPut)
	router.HandleFunc("/accounts/{account_id}", middleware.MakeHTTPHandleFunc(s.deleteAccountById)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.AccountPort, router)
	if err != nil {
		return
	}
}
