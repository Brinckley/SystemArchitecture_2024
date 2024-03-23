package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user_service/internal/util"
)

func (s *UserApiServer) createAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl)
	if err != nil {
		return err
	}

	log.Println("Sending data for creation to the account service")
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, http.StatusOK, accountResp)
}

func (s *UserApiServer) getAccounts(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl)
	if err != nil {
		return err
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, http.StatusOK, accountResp)
}

func (s *UserApiServer) getAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	id := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/"+id)
	if err != nil {
		return err
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, http.StatusOK, accountResp)
}

func (s *UserApiServer) updateAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	id := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/"+id)
	if err != nil {
		return err
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, http.StatusOK, accountResp)
}

func (s *UserApiServer) deleteAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	id := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/"+id)
	if err != nil {
		return err
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return err
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, http.StatusOK, accountResp)
}
