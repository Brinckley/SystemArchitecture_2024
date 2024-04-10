package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"post_service/internal"
)

func (s *PostApiServer) createPost(w http.ResponseWriter, r *http.Request) error {
	var createPostReq internal.PostDto
	if err := json.NewDecoder(r.Body).Decode(&createPostReq); err != nil {
		return writeJson(w, http.StatusBadRequest, fmt.Sprintf("fail to handle data error %v", err))
	}
	postId, err := s.Storage.Create(*s.Context, createPostReq)
	if err != nil {
		return writeJson(w, http.StatusNoContent, err)
	}
	return writeJson(w, http.StatusOK, fmt.Sprintf("Id of the created post %s", postId))
}

func (s *PostApiServer) getPost(w http.ResponseWriter, r *http.Request) error {
	postId := mux.Vars(r)["post_id"]
	postById, err := s.Storage.GetById(*s.Context, postId)
	if err != nil {
		return writeJson(w, http.StatusNoContent, err)
	}
	return writeJson(w, http.StatusOK, postById)
}

func (s *PostApiServer) getPostsByAccId(w http.ResponseWriter, r *http.Request) error {
	accountId := mux.Vars(r)["account_id"]
	posts, err := s.Storage.GetByAccountId(*s.Context, accountId)
	if err != nil {
		return writeJson(w, http.StatusNoContent, err)
	}
	return writeJson(w, http.StatusOK, posts)
}

func (s *PostApiServer) updatePost(w http.ResponseWriter, r *http.Request) error {
	postId := mux.Vars(r)["post_id"]
	var updatePostReq internal.Post
	if err := json.NewDecoder(r.Body).Decode(&updatePostReq); err != nil {
		return writeJson(w, http.StatusBadRequest, fmt.Sprintf("fail to handle data error %v", err))
	}
	updatePostReq.Id = postId
	err := s.Storage.Update(*s.Context, updatePostReq)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, err)
	}
	return writeJson(w, http.StatusOK, "post updated")
}

func (s *PostApiServer) deletePost(w http.ResponseWriter, r *http.Request) error {
	postId := mux.Vars(r)["post_id"]
	err := s.Storage.Delete(*s.Context, postId)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, err)
	}
	return writeJson(w, http.StatusOK, "post deleted")
}
