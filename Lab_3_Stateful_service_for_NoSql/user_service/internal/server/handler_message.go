package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/util"
)

func (s *UserApiServer) createMessage(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl)
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

func (s *UserApiServer) getMessages(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/account/"+accountId)
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

func (s *UserApiServer) getMessage(responseWriter http.ResponseWriter, userReq *http.Request) error {
	messageId := mux.Vars(userReq)["message_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/"+messageId)
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
