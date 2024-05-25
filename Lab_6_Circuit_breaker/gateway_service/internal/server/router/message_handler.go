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
	UNABLE_TO_SEND_MSG_PROXY_REQ   = "unable to send proxy request for messages"
	UNABLE_TO_CREATE_MSG_PROXY_REQ = "unable to create proxy request for messages"
)

func (s *UserApiServer) createMessage(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/"+accountId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_MSG_PROXY_REQ)
	}

	msgResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_MSG_PROXY_REQ)
	}
	util.CopyHeadersToWriter(msgResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, msgResp.StatusCode, msgResp)
}

func (s *UserApiServer) getMessage(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	messageId := mux.Vars(userReq)["message_id"]

	messageFromCache, err := s.Cache.GetMessage(messageId)
	if err != nil {
		var cacheError *storage.CacheError
		ok := errors.As(err, &cacheError)
		if !ok {
			log.Printf("[ERR] Cache error %s", err)
		} else {
			log.Printf("[ERR] Failed to get message from cache for message id %s, error : %s", messageId, err)
		}
	} else {
		return middleware.WriteJson(responseWriter, http.StatusOK, messageFromCache)
	}

	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/"+accountId+"/msg/"+"/"+messageId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_MSG_PROXY_REQ)
	}
	msgResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_MSG_PROXY_REQ)
	}

	util.CopyHeadersToWriter(msgResp, responseWriter)
	if msgResp.StatusCode == http.StatusOK {
		messageToCache, err := s.writeMessageToCache(messageId, msgResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
			return middleware.WriteJson(responseWriter, msgResp.StatusCode, messageToCache)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, messageToCache)
	}
	return middleware.WriteJsonFromResponse(responseWriter, msgResp.StatusCode, msgResp)
}

func (s *UserApiServer) getMessages(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/account/"+accountId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_MSG_PROXY_REQ)
	}

	msgResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_MSG_PROXY_REQ)
	}

	// TODO: Circuit Breaker

	util.CopyHeadersToWriter(msgResp, responseWriter)
	if msgResp.StatusCode == http.StatusOK {
		messagesToCache, err := s.writeMessagesToCache(accountId, msgResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
			return middleware.WriteJson(responseWriter, msgResp.StatusCode, messagesToCache)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, messagesToCache)
	}
	return middleware.WriteJsonFromResponse(responseWriter, msgResp.StatusCode, msgResp)
}

func (s *UserApiServer) writeMessageToCache(messageId string, body io.ReadCloser) (message entity.Message, err error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return message, err
	}
	err = json.Unmarshal(bytes, &message)
	if err != nil {
		return message, err
	}
	return message, s.Cache.SetMessage(messageId, message)
}

func (s *UserApiServer) writeMessagesToCache(accountId string, body io.ReadCloser) (messages entity.MessagesCollection, err error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return messages, err
	}
	err = json.Unmarshal(bytes, &messages)
	if err != nil {
		return messages, err
	}
	err = s.Cache.SetMessageCollection(accountId, messages)
	if err != nil {
		return messages, err
	}
	return messages, nil
}
