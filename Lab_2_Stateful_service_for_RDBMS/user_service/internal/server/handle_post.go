package server

import "net/http"

func (s *UserApiServer) createPost(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) getPosts(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) getPost(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) updatePost(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}

func (s *UserApiServer) deletePost(w http.ResponseWriter, r *http.Request) error {

	return writeJson(w, http.StatusOK, []byte{})
}
