package server

import "net/http"

func (s *UserApiServer) createMessage(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) getMessages(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) getMessage(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) updateMessage(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) deleteMessage(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}
