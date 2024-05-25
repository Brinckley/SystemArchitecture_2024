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
	"strconv"
)

const (
	UNABLE_TO_SEND_ACCOUNT_PROXY_REQ   = "unable to send proxy request for accounts"
	UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ = "unable to create proxy request for accounts"
)

func (s *UserApiServer) getAccounts(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts")
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_ACCOUNT_PROXY_REQ)
	}

	util.CopyHeadersToWriter(accountResp, responseWriter)
	if accountResp.StatusCode == http.StatusOK {
		accountsToCache, err := s.writeAccountsToCache(accountResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
			return middleware.WriteJson(responseWriter, accountResp.StatusCode, accountsToCache)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, accountsToCache)
	}
	return middleware.WriteJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) getAccount(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	accountId := mux.Vars(userReq)["account_id"]

	accountCache, err := s.Cache.GetAccount(accountId)
	if err != nil {
		var cacheError *storage.CacheError
		ok := errors.As(err, &cacheError)
		if !ok {
			log.Printf("[ERR] Cache error %s", err)
		} else {
			log.Printf("[ERR] Failed to get account from cache for account id %s, error : %s", accountId, err)
		}
	} else {
		return middleware.WriteJson(responseWriter, http.StatusOK, accountCache)
	}

	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts/"+accountId)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ)
	}

	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_ACCOUNT_PROXY_REQ)
	}

	util.CopyHeadersToWriter(accountResp, responseWriter)
	if accountResp.StatusCode == http.StatusOK {
		accountToCache, err := s.writeAccountToCache(accountId, accountResp.Body)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
			return middleware.WriteJson(responseWriter, accountResp.StatusCode, accountToCache)
		}
		return middleware.WriteJson(responseWriter, http.StatusOK, accountToCache)
	}
	return middleware.WriteJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) updateAccount(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts/"+accountId)
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

func (s *UserApiServer) deleteAccount(responseWriter http.ResponseWriter, userReq *http.Request, accountId string) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/accounts/"+accountId)
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

func (s *UserApiServer) getAccountsByMask(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/search")
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

func (s *UserApiServer) writeAccountToCache(accountId string, body io.ReadCloser) (account entity.Account, err error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return account, err
	}
	err = json.Unmarshal(bytes, &account)
	if err != nil {
		return account, err
	}
	return account, s.Cache.SetAccount(accountId, account)
}

func (s *UserApiServer) writeAccountsToCache(body io.ReadCloser) (accounts entity.AccountCollection, err error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return accounts, err
	}
	err = json.Unmarshal(bytes, &accounts)
	if err != nil {
		return accounts, err
	}

	for _, account := range accounts.Accounts {
		err := s.Cache.SetAccount(strconv.Itoa(account.Id), account)
		if err != nil {
			log.Printf("[ERR] Failed to write account to cache error %s", err)
		}
	}

	return accounts, nil
}
