package router

import (
	"Gateway/internal/server/middleware"
	"Gateway/internal/server/response_error"
	"Gateway/internal/server/util"
	"net/http"
)

func (s *UserApiServer) signUpAccount(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/signup")
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_ACCOUNT_PROXY_REQ)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) signInAccount(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/signin")
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_ACCOUNT_PROXY_REQ)
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}
