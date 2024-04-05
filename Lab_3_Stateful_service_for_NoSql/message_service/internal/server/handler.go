package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"message_service/internal"
	"net/http"
)

func (s *MessageApiServer) getAllMessagesTo(w http.ResponseWriter, r *http.Request) error {
	accountId := mux.Vars(r)["account_id"]
	messages, err := s.Storage.GetByDestId(*s.Context, accountId)
	if err != nil {
		return writeJson(w, http.StatusNoContent, err)
	}
	return writeJson(w, http.StatusOK, messages)
}

func (s *MessageApiServer) createMessage(w http.ResponseWriter, r *http.Request) error {
	var msgDto internal.MessageDto
	if err := json.NewDecoder(r.Body).Decode(&msgDto); err != nil {
		return writeJson(w, http.StatusBadRequest, fmt.Errorf("fail to handle request body error %v", err))
	}
	messageId, err := s.Storage.Create(*s.Context, msgDto)
	if err != nil {
		return writeJson(w, http.StatusNoContent, err)
	}
	return writeJson(w, http.StatusOK, fmt.Sprintf("msg with id %s created", messageId))
}

func (s *MessageApiServer) getMessageByDestId(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	messages, err := s.Storage.GetByDestId(*s.Context, accountIdRaw)
	if err != nil {
		return writeJson(w, http.StatusNoContent, err)
	}
	return writeJson(w, http.StatusOK, messages)
}
