package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/server/middleware"
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
	router.Use(middleware.LoggingMiddleWare)

	router.HandleFunc("/signup", middleware.MakeHTTPHandleFunc(s.signUpAccount)).Methods(http.MethodPost)
	router.HandleFunc("/signup", middleware.MakeHTTPHandleFunc(s.signInAccount)).Methods(http.MethodGet)

	router.HandleFunc("/account", middleware.MakeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)
	router.HandleFunc("/account/search", middleware.MakeHTTPHandleFunc(s.getAccountsByMask)).Methods(http.MethodGet)

	router.HandleFunc("/account/{account_id}", middleware.MakeHTTPHandleFunc(s.getAccount)).Methods(http.MethodGet)
	router.HandleFunc("/account/{account_id}", middleware.MakeHTTPHandleFunc(s.updateAccount)).Methods(http.MethodPut)
	router.HandleFunc("/account/{account_id}", middleware.MakeHTTPHandleFunc(s.deleteAccount)).Methods(http.MethodDelete)

	router.HandleFunc("/messages/{message_id}", middleware.MakeHTTPHandleFunc(s.getMessage)).Methods(http.MethodGet)
	router.HandleFunc("/messages", middleware.MakeHTTPHandleFunc(s.createMessage)).Methods(http.MethodPost)
	router.HandleFunc("/messages/account/{account_id}", middleware.MakeHTTPHandleFunc(s.getMessages)).Methods(http.MethodGet)

	router.HandleFunc("/posts", middleware.MakeHTTPHandleFunc(s.createPost)).Methods(http.MethodPost)
	router.HandleFunc("/posts/account/{account_id}", middleware.MakeHTTPHandleFunc(s.getPosts)).Methods(http.MethodGet)
	router.HandleFunc("/posts/{post_id}", middleware.MakeHTTPHandleFunc(s.getPost)).Methods(http.MethodGet)
	router.HandleFunc("/posts/{post_id}", middleware.MakeHTTPHandleFunc(s.updatePost)).Methods(http.MethodPut)
	router.HandleFunc("/posts/{post_id}", middleware.MakeHTTPHandleFunc(s.deletePost)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.UserPort, router)
	if err != nil {
		return err
	}
	return nil
}
