package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"post_service/internal"
	"strconv"
)

func (s *PostApiServer) getPosts(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	posts, err := s.Storage.GetPostsByAccountId(accountId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, posts)
}

func (s *PostApiServer) createPost(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	var createPostReq internal.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&createPostReq); err != nil {
		return err
	}
	createPostReq.AccountId = accountId
	log.Println(createPostReq)
	postId, err := s.Storage.CreatePost(&createPostReq)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, postId)
}

func (s *PostApiServer) getPost(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	postIdRaw := mux.Vars(r)["post_id"]
	postId, err := strconv.Atoi(postIdRaw)

	var modPost internal.CreatePostRequest
	err = json.NewDecoder(r.Body).Decode(&modPost)
	if err != nil {
		return err
	}
	account := internal.PostFrom(idInt, &modAccount)

	postById, err := s.Storage.GetPostByAccountById(accountId, postId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, postById)
}

func (s *PostApiServer) updatePost(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	postIdRaw := mux.Vars(r)["post_id"]
	postId, err := strconv.Atoi(postIdRaw)
	modifiedPost, err := s.Storage.UpdatePostByAccountById(accountId, postId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, modifiedPost)
}

func (s *PostApiServer) deletePost(w http.ResponseWriter, r *http.Request) error {
	accountIdRaw := mux.Vars(r)["account_id"]
	accountId, err := strconv.Atoi(accountIdRaw)
	postIdRaw := mux.Vars(r)["post_id"]
	postId, err := strconv.Atoi(postIdRaw)
	deletedPostId, err := s.Storage.DeletePostByAccountById(accountId, postId)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, deletedPostId)
}
