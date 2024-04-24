package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/server/middleware"
	"user_service/internal/server/response_error"
	"user_service/internal/util"
)

const (
	UNABLE_TO_SEND_POSTS_PROXY_REQ   = "unable to send proxy request for posts"
	UNABLE_TO_CREATE_POSTS_PROXY_REQ = "unable to create proxy request for posts"
)

func (s *UserApiServer) createPost(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+accountId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_POSTS_PROXY_REQ)
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_POSTS_PROXY_REQ)
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getPosts(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	accountId := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_POSTS_PROXY_REQ)
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_POSTS_PROXY_REQ)
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getPost(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+postId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_POSTS_PROXY_REQ)
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_POSTS_PROXY_REQ)
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) updatePost(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+accountId+"/posts/"+postId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_POSTS_PROXY_REQ)
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_POSTS_PROXY_REQ)
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) deletePost(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	postId := mux.Vars(userReq)["post_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+accountId+"/posts/"+postId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_POSTS_PROXY_REQ)
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_POSTS_PROXY_REQ)
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}
