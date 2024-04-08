package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type UserApiServer struct {
	UserPort   string
	AccountUrl string
	MsgUrl     string
	PostUrl    string
}

func NewUserApiServer(userPort, accountUrl, msgUrl, postUrl string) *UserApiServer {
	return &UserApiServer{
		UserPort:   userPort,
		AccountUrl: accountUrl,
		MsgUrl:     msgUrl,
		PostUrl:    postUrl,
	}
}

func (s *UserApiServer) Run() error {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	router.HandleFunc("/account", makeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)
	router.HandleFunc("/account", makeHTTPHandleFunc(s.createAccount)).Methods(http.MethodPost)
	router.HandleFunc("/account/search", makeHTTPHandleFunc(s.getAccountsByMask)).Methods(http.MethodGet)
	router.HandleFunc("/account/{account_id}", makeHTTPHandleFunc(s.getAccount)).Methods(http.MethodGet)
	router.HandleFunc("/account/{account_id}", makeHTTPHandleFunc(s.updateAccount)).Methods(http.MethodPut)
	router.HandleFunc("/account/{account_id}", makeHTTPHandleFunc(s.deleteAccount)).Methods(http.MethodDelete)

	router.HandleFunc("/messages/{message_id}", makeHTTPHandleFunc(s.getMessage)).Methods(http.MethodGet)
	router.HandleFunc("/messages", makeHTTPHandleFunc(s.createMessage)).Methods(http.MethodPost)
	router.HandleFunc("/messages/account/{account_id}", makeHTTPHandleFunc(s.getMessages)).Methods(http.MethodGet)

	router.HandleFunc("/posts", makeHTTPHandleFunc(s.createPost)).Methods(http.MethodPost)
	router.HandleFunc("/posts/account/{account_id}", makeHTTPHandleFunc(s.getPosts)).Methods(http.MethodGet)
	router.HandleFunc("/posts/{post_id}", makeHTTPHandleFunc(s.getPost)).Methods(http.MethodGet)
	router.HandleFunc("/posts/{post_id}", makeHTTPHandleFunc(s.updatePost)).Methods(http.MethodPut)
	router.HandleFunc("/posts/{post_id}", makeHTTPHandleFunc(s.deletePost)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.UserPort, router)
	if err != nil {
		return err
	}
	return nil
}
