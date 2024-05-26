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
	if err != nil { // smth went wrong -> need to update circuit breaker
		err := s.CircuitBreakerMessage.UpdateState() // incrementing
		if err != nil {
			return middleware.WriteJson(responseWriter, http.StatusInternalServerError,
				fmt.Sprintf("something fatally went wrong %s", err))
		}
		if s.CircuitBreakerMessage.IsCircuitOpen() { // checking if counter has achieved the max mark
			messages, err := s.Cache.GetMessageCollection(accountId)
			if errors.Is(err, redis.Nil) { // data's gone down by TTL, or it has never existed
				return middleware.WriteJson(responseWriter, http.StatusInternalServerError,
					"Nothing in the cache anymore. App is dead. Please, wait.")
			} else if err != nil {
				return middleware.WriteJson(responseWriter, http.StatusInternalServerError,
					fmt.Sprintf("something fatally went wrong with fetching data from cache %s", err))
			}
			return middleware.WriteJson(responseWriter, http.StatusOK, messages.Messages)
		}
		return middleware.WriteJson(responseWriter, http.StatusInternalServerError, "Something went wrong, try again later.")
	}

	util.CopyHeadersToWriter(msgResp, responseWriter)
	statusCode := msgResp.StatusCode
	log.Println("MessageService statusCode:", statusCode)
	if statusCode == http.StatusOK {
		err := s.CircuitBreakerMessage.ClearCounter() // if everything is ok we clear the counter
		if err != nil {
			log.Printf("[ERR] Failed to write clear the circuitBreaker Message counter %s", err)
		}
		messagesToCache, err := s.writeMessagesToCache(accountId, msgResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write messages to cache error %s", err)
			return middleware.WriteJson(responseWriter, statusCode, messagesToCache.Messages)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, messagesToCache.Messages)
	}
	return middleware.WriteJsonFromResponse(responseWriter, statusCode, msgResp)
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
	err = json.Unmarshal(bytes, &messages.Messages)
	if err != nil {
		return messages, err
	}

	err = s.Cache.SetMessageCollection(accountId, messages)
	if err != nil {
		return messages, err
	}
	return messages, nil
}
