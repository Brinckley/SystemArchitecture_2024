package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"user_service/internal/util"
)

func (s *UserApiServer) createAccount(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts")
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

func (s *UserApiServer) getAccounts(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts")
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
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts/"+id)
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
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts/"+id)
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
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts/"+id)
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

func (s *UserApiServer) getAccountsByMask(responseWriter http.ResponseWriter, userReq *http.Request) error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/search")
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
