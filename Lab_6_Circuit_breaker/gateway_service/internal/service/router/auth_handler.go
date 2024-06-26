package router

import (
	"Gateway/internal/entity"
	"Gateway/internal/service/middleware"
	"Gateway/internal/service/response_error"
	"Gateway/internal/service/util"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	ERR_DECODE_DATA     = "cannot decode account data"
	ERR_FAILED_ID_CACHE = "Failed to work with ID"
)

func copyReq(req *http.Request) (*http.Request, error) {
	var b bytes.Buffer
	_, err := b.ReadFrom(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(&b)
	var copyReq http.Request
	copyReq.Body = io.NopCloser(bytes.NewReader(b.Bytes()))
	return &copyReq, nil
}

type AccountId struct {
	AccountId int `json:"account_id"`
}

func (s *UserApiServer) signUpAccount(responseWriter http.ResponseWriter, userReq *http.Request) *response_error.Error {
	copyReq, err := copyReq(userReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ)
	}
	proxyReq, err := util.CreateProxyRequest(userReq, s.AccountUrl+"/signup")
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_CREATE_ACCOUNT_PROXY_REQ)
	}
	accountResp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, UNABLE_TO_SEND_ACCOUNT_PROXY_REQ)
	}

	if accountResp.StatusCode == http.StatusOK {
		var signUpAccount entity.SignUpAccount
		if err := json.NewDecoder(copyReq.Body).Decode(&signUpAccount); err != nil {
			return middleware.WriteJson(responseWriter, http.StatusInternalServerError, ERR_DECODE_DATA)
		}
		bodyBytes, err := io.ReadAll(accountResp.Body)
		if err != nil {
			return middleware.WriteJson(responseWriter, http.StatusInternalServerError, ERR_FAILED_ID_CACHE)
		}
		rawId := util.ConvertToStringWithoutBlanks(bodyBytes)
		intId, err := strconv.Atoi(rawId)
		if err != nil {
			return middleware.WriteJson(responseWriter, http.StatusInternalServerError, ERR_FAILED_ID_CACHE)
		}
		err = s.writeAccountWithIdToCache(signUpAccount, intId)
		if err != nil {
			log.Println("unable to write account to cache")
		}
		responseWriter.WriteHeader(http.StatusOK)
		err = json.NewEncoder(responseWriter).Encode(AccountId{AccountId: intId})
		if err != nil {
			return response_error.New(err, http.StatusInternalServerError, "cannot encode error")
		}
		return nil
	}
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}

func (s *UserApiServer) writeAccountWithIdToCache(accountSignUp entity.SignUpAccount, id int) error {
	account := entity.Account{
		Id:        id,
		Username:  accountSignUp.Username,
		Password:  accountSignUp.Password,
		FirstName: accountSignUp.FirstName,
		LastName:  accountSignUp.LastName,
		Email:     accountSignUp.Email,
	}
	return s.Cache.SetAccount(strconv.Itoa(id), account)
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
	log.Printf("Header : %s", accountResp.Header.Get("Auth-token"))
	util.CopyHeadersToWriter(accountResp, responseWriter)
	return middleware.WriteJsonFromResponse(responseWriter, accountResp.StatusCode, accountResp)
}
