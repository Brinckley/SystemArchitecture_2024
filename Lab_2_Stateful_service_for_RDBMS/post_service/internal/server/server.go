package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"post_service/internal/repository"
)

type PostApiServer struct {
	Storage  repository.Storage
	PostPort string
}

func NewPostApiServer(port string, storage repository.Storage) *PostApiServer {
	return &PostApiServer{
		Storage:  storage,
		PostPort: port,
	}
}

func (s *PostApiServer) Run() {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	router.HandleFunc("/{account_id}", makeHTTPHandleFunc(s.getPosts)).Methods(http.MethodGet)
	router.HandleFunc("/{account_id}", makeHTTPHandleFunc(s.createPost)).Methods(http.MethodPost)
	router.HandleFunc("/{account_id}/{post_id}", makeHTTPHandleFunc(s.getPost)).Methods(http.MethodGet)
	router.HandleFunc("/{account_id}/{post_id}", makeHTTPHandleFunc(s.updatePost)).Methods(http.MethodPut)
	router.HandleFunc("/{account_id}/{post_id}", makeHTTPHandleFunc(s.deletePost)).Methods(http.MethodDelete)

	err := http.ListenAndServe(":"+s.PostPort, router)
	if err != nil {
		return
	}
}
