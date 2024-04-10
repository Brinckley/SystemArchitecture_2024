package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/util"
)

func (s *UserApiServer) createPost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl)
	if err != nil {
		return err
	}

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

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getPost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+postId)
	if err != nil {
		return err
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) updatePost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+postId)
	if err != nil {
		return err
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) deletePost(responseWriter http.ResponseWriter, userReq *http.Request) error {
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+postId)
	if err != nil {
		return err
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}
