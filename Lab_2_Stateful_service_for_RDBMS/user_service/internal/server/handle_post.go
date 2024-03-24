package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user_service/internal/util"
)

func (s *UserApiServer) createPost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the post service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getPosts(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the post service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getPost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId+"/"+postId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the post service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) updatePost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId+"/"+postId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the post service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) deletePost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId+"/"+postId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the post service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}
