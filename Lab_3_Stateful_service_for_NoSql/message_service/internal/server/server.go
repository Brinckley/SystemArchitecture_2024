package server

import (
	"github.com/gorilla/mux"
	"message_service/internal/repository"
	"net/http"
)

type MessageApiServer struct {
	Storage     repository.Storage
	MessagePort string
}

func NewMessageApiServer(port string, storage repository.Storage) *MessageApiServer {
	return &MessageApiServer{
		Storage:     storage,
		MessagePort: port,
	}
}

func (s *MessageApiServer) Run() {
	router := mux.NewRouter()
	router.Use(loggingMiddleWare)

	// getting all messages that the user with accountId has received
	router.HandleFunc("/{account_id}", makeHTTPHandleFunc(s.getAllMessagesTo)).Methods(http.MethodGet)
	// creating new message from account id as sender
	router.HandleFunc("/{account_id}", makeHTTPHandleFunc(s.createMessage)).Methods(http.MethodPost)
	// getting message that the user with accountId has received with id
	router.HandleFunc("/{account_id}/{id}", makeHTTPHandleFunc(s.getMessageToById)).Methods(http.MethodGet)

	err := http.ListenAndServe(":"+s.MessagePort, router)
	if err != nil {
		return
	}
}
