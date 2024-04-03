package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user_service/internal/util"
)

func (s *UserApiServer) createMessage(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/"+accountId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the message service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getMessages(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.MsgUrl+"/"+accountId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the message service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}

func (s *UserApiServer) getMessage(responseWriter http.ResponseWriter, userReq *http.Request) error {
	accountId := mux.Vars(userReq)["account_id"]
	msgId := mux.Vars(userReq)["msg_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.PostUrl+"/"+accountId+"/"+msgId)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the message service")
	postResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(postResp, responseWriter)
	return writeJsonFromResponse(responseWriter, postResp.StatusCode, postResp)
}
