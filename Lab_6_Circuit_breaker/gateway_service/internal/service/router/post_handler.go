package router

import (
	"Gateway/internal/entity"
	"Gateway/internal/service/middleware"
	"Gateway/internal/service/response_error"
	"Gateway/internal/service/util"
	"Gateway/internal/storage"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
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
	if err != nil { // smth went wrong -> need to update circuit breaker
		err := s.CircuitBreakerPost.UpdateState() // incrementing
		if err != nil {
			return middleware.WriteJson(responseWriter, http.StatusInternalServerError,
				fmt.Sprintf("something fatally went wrong %s", err))
		}
		if s.CircuitBreakerPost.IsCircuitOpen() { // checking if counter has achieved the max mark
			log.Printf("[INFO] getting all posts for account %s from cache", accountId)
			posts, err := s.Cache.GetPostCollection(accountId)
			if errors.Is(err, redis.Nil) { // data's gone down by TTL, or it has never existed
				return middleware.WriteJson(responseWriter, http.StatusInternalServerError,
					"Nothing in the cache anymore. App is dead. Please, wait.")
			} else if err != nil {
				return middleware.WriteJson(responseWriter, http.StatusInternalServerError,
					fmt.Sprintf("something fatally went wrong with fetching data from cache %s", err))
			}
			return middleware.WriteJson(responseWriter, http.StatusOK, posts)
		}
		return middleware.WriteJson(responseWriter, http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	util.CopyHeadersToWriter(postResp, responseWriter)
	statusCode := postResp.StatusCode
	log.Println("PostService statusCode:", statusCode)
	if statusCode == http.StatusOK {
		err := s.CircuitBreakerPost.ClearCounter()
		if err != nil {
			log.Printf("[ERR] Failed to write clear the circuitBreaker Post counter %s", err)
		}
		postsToCache, err := s.writePostsToCache(accountId, postResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write posts to cache error %s", err)
			return middleware.WriteJson(responseWriter, statusCode, postsToCache.Posts)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, postsToCache.Posts)
	}
	return middleware.WriteJsonFromResponse(responseWriter, statusCode, postResp)
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
	if postResp.StatusCode == http.StatusOK {
		err := s.Cache.DeletePost(postId)
		if err != nil {
			log.Printf("[ERR] Failed to delete post with id %s for account %s from cache error %s", postId, accountId, err)
		}
	}
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
	if postResp.StatusCode == http.StatusOK {
		err := s.Cache.DeletePost(postId)
		if err != nil {
			log.Printf("[ERR] Failed to delete post with id %s for account %s from cache error %s", postId, accountId, err)
		}
	}
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
	err = json.Unmarshal(bytes, &posts.Posts)
	if err != nil {
		return posts, err
	}

	err = s.Cache.SetPostCollection(accountId, posts)
	if err != nil {
		return posts, err
	}
	return posts, nil
}
