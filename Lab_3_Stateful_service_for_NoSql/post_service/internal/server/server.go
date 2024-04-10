package server

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"post_service/internal/storage"
)

type PostApiServer struct {
	Storage  storage.Storage
	PostPort string
	Context  *context.Context
}

func NewPostApiServer(port string, storage storage.Storage, ctx *context.Context) *PostApiServer {
	return &PostApiServer{
		Storage:  storage,
		PostPort: port,
		Context:  ctx,
	}
}

func (s *PostApiServer) Run() {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	router.HandleFunc("/", makeHTTPHandleFunc(s.createPost)).Methods(http.MethodPost)
	router.HandleFunc("/{account_id}", makeHTTPHandleFunc(s.getPostsByAccId)).Methods(http.MethodGet)
	router.HandleFunc("/account/{post_id}", makeHTTPHandleFunc(s.getPost)).Methods(http.MethodGet)
	router.HandleFunc("/account/{post_id}", makeHTTPHandleFunc(s.updatePost)).Methods(http.MethodPut)
	router.HandleFunc("/account/{post_id}", makeHTTPHandleFunc(s.deletePost)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.PostPort, router)
	if err != nil {
		return
	}
}
