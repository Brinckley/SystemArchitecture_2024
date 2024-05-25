package router

import (
	"Gateway/internal/entity"
	"Gateway/internal/server/middleware"
	"Gateway/internal/server/response_error"
	"Gateway/internal/server/util"
	"Gateway/internal/storage"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
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

	// TODO: Circuit Breaker

	util.CopyHeadersToWriter(postResp, responseWriter)
	if postResp.StatusCode == http.StatusOK {
		postsToCache, err := s.writePostsToCache(accountId, postResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
			return middleware.WriteJson(responseWriter, postResp.StatusCode, postsToCache)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, postsToCache)
	}
	return middleware.WriteJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getPost(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	postId := mux.Vars(userReq)["post_id"]

	postFromCache, err := s.Cache.GetPost(postId)
	if err != nil {
		var cacheError *storage.CacheError
		ok := errors.As(err, &cacheError)
		if !ok {
			log.Printf("[ERR] Cache error %s", err)
		} else {
			log.Printf("[ERR] Failed to get account from cache for post id %s, error : %s", postId, err)
		}
	} else {
		return middleware.WriteJson(responseWriter, http.StatusOK, postFromCache)
	}

	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/account/"+postId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_POSTS_PROXY_REQ)
	}

	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_POSTS_PROXY_REQ)
	}

	util.CopyHeadersToWriter(postResp, responseWriter)
	if postResp.StatusCode == http.StatusOK {
		postToCache, err := s.writePostToCache(postId, postResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
			return middleware.WriteJson(responseWriter, postResp.StatusCode, postToCache)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, postToCache)
	}
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

func (s *UserApiServer) writePostToCache(postId string, body io.ReadCloser) (post entity.Post, err error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return post, err
	}
	err = json.Unmarshal(bytes, &post)
	if err != nil {
		return post, err
	}
	return post, s.Cache.SetPost(postId, post)
}

func (s *UserApiServer) writePostsToCache(accountId string, body io.ReadCloser) (posts entity.PostCollection, err error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return posts, err
	}
	err = json.Unmarshal(bytes, &posts)
	if err != nil {
		return posts, err
	}

	return accounts, nil
}
