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
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}

	log.Println("Sending data for creation to the account service")
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) getAccounts(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) getAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	id := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/"+id)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) updateAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	id := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/"+id)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) deleteAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	id := mux.Vars(userReq)["account_id"]
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/"+id)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return writeJson(responseWriter, http.StatusInternalServerError, err)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return writeJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}
