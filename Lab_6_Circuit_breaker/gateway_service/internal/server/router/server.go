package router

import (
	"Gateway/internal/server/middleware"
	"Gateway/internal/storage/cache"
	"github.com/gorilla/mux"
	"net/http"
)

type UserApiServer struct {
	UserPort   string
	AccountUrl string
	MsgUrl     string
	PostUrl    string
	Cache      cache.SNCache
}

func NewUserApiServer(userPort, accountUrl, msgUrl, postUrl string, cache cache.SNCache) *UserApiServer {
	return &UserApiServer{
		UserPort:   userPort,
		AccountUrl: accountUrl,
		MsgUrl:     msgUrl,
		PostUrl:    postUrl,
		Cache:      cache,
	}
}

func (s *UserApiServer) Run() error {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleWare)

	// no need auth block
	router.HandleFunc("/signup", middleware.MakeHTTPHandleFunc(s.signUpAccount)).Methods(http.MethodPost)
	router.HandleFunc("/signin", middleware.MakeHTTPHandleFunc(s.signInAccount)).Methods(http.MethodPost)

	// no need auth block
	router.HandleFunc("/accounts", middleware.MakeHTTPHandleFunc(s.getAccounts)).Methods(http.MethodGet)
	router.HandleFunc("/accounts/search", middleware.MakeHTTPHandleFunc(s.getAccountsByMask)).Methods(http.MethodGet)

	// no need auth block
	router.HandleFunc("/account/{account_id}", middleware.MakeHTTPHandleFunc(s.getAccount)).Methods(http.MethodGet)

	// need auth block
	router.HandleFunc("/account", middleware.MakeHTTPAuthedHandleFunc(s.updateAccount, s.AccountUrl)).Methods(http.MethodPut)
	router.HandleFunc("/account", middleware.MakeHTTPAuthedHandleFunc(s.deleteAccount, s.AccountUrl)).Methods(http.MethodDelete)

	// need auth block
	router.HandleFunc("/messages/msg/{message_id}", middleware.MakeHTTPAuthedHandleFunc(s.getMessage, s.AccountUrl)).Methods(http.MethodGet)
	router.HandleFunc("/messages/account", middleware.MakeHTTPAuthedHandleFunc(s.getMessages, s.AccountUrl)).Methods(http.MethodGet)
	router.HandleFunc("/messages", middleware.MakeHTTPAuthedHandleFunc(s.createMessage, s.AccountUrl)).Methods(http.MethodPost)

	// no need auth block
	router.HandleFunc("/posts/account/{account_id}", middleware.MakeHTTPHandleFunc(s.getPosts)).Methods(http.MethodGet)
	router.HandleFunc("/posts/{post_id}", middleware.MakeHTTPHandleFunc(s.getPost)).Methods(http.MethodGet)

	// need auth block
	router.HandleFunc("/posts", middleware.MakeHTTPAuthedHandleFunc(s.createPost, s.AccountUrl)).Methods(http.MethodPost)
	router.HandleFunc("/posts/{post_id}", middleware.MakeHTTPAuthedHandleFunc(s.updatePost, s.AccountUrl)).Methods(http.MethodPut)
	router.HandleFunc("/posts/{post_id}", middleware.MakeHTTPAuthedHandleFunc(s.deletePost, s.AccountUrl)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.UserPort, router)
	if err != nil {
		return err
	}
	return nil
}
