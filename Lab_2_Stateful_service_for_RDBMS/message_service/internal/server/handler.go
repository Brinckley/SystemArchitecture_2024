package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"message_service/internal"
	"net/http"
	"strconv"
)

func (s *MessageApiServer) getAllMessagesTo(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	if err != nil {
		return err
	}
	messages, err := s.Storage.GetMessagesByReceiverId(accountId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, messages)
}

func (s *MessageApiServer) createMessage(w http.ResponseWriter, r *http.Request) error {
	var createMessageReq internal.CreateMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&createMessageReq); err != nil {
		return err
	}
	messageId, err := s.Storage.CreateMessage(&createMessageReq)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, messageId)
}

func (s *MessageApiServer) getMessageToById(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	if err != nil {
		return err
	}
	msgIdRaw := mux.Vars(r)["id"]
	msgId, err := strconv.Atoi(msgIdRaw)
	if err != nil {
		return err
	}
	id, err := s.Storage.GetMessageById(accountId, msgId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, id)
}
