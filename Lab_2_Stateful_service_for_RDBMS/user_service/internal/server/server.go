package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ApiServer struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewApiServer(host, port string) *ApiServer {
	return &ApiServer{
		Host: host,
		Port: port,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	http.ListenAndServe("", router)
}
