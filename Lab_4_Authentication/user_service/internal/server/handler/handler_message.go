package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/server/middleware"
	"user_service/internal/server/response_error"
	"user_service/internal/util"
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
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/"+accountId+"/msg/"+"/"+messageId)
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

func (s *UserApiServer) getMessages(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/account/"+accountId)
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
