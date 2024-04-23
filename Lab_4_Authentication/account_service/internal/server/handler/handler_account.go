package handler

import (
	"account_service/internal"
	"account_service/internal/server/middleware"
	"account_service/internal/server/response_error"
	"account_service/internal/util"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

const (
	signingKey                = "wi2omp{wr2o*#3he3uJls@hsl3sa48"
	MSG_NO_ACCOUNTS           = "no accounts found"
	MSG_SUCCESS               = "operation successful"
	ERR_DECODE_DATA           = "cannot decode account data"
	ERR_PASSWORD_GENERATION   = "cannot create password for the account"
	ERR_PASSWORD_MATCH        = "cannot match password for the account"
	ERR_SIGN_UP               = "cannot create account or already exists"
	ERR_GET_ALL               = "cannot get all accounts"
	ERR_GET_FIND_ACCOUNTS     = "cannot find/get matched accounts"
	ERR_CANNOT_DECODE_ID      = "cannot decode account's id"
	ERR_CANNOT_GET_ACCOUNT    = "cannot find/get the account"
	ERR_CANNOT_GET_ACCOUNT_ID = "cannot find/get the account with such id"
	ERR_SIGN_IN_USERNAME      = "cannot sign in the account. check login"
	ERR_SIGN_IN_PASSWORD      = "cannot sign in the account. check password"
	ERR_HANDLE_TOKEN          = "cannot create token"
)

func (s *AccountApiServer) signUpAccount(w http.ResponseWriter, r *http.Request) *response_error.Error {
	var signUpAccount internal.SignUpAccount
	if err := json.NewDecoder(r.Body).Decode(&signUpAccount); err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_DECODE_DATA)
	}
	newPassword, err := util.HashPassword(signUpAccount.Password)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_PASSWORD_GENERATION)
	}
	signUpAccount.Password = newPassword
	accountId, err := s.Storage.SignUpAccount(signUpAccount)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_SIGN_UP)
	}
	return middleware.WriteJson(w, http.StatusOK, accountId)
}

func (s *AccountApiServer) signInAccount(w http.ResponseWriter, r *http.Request) *response_error.Error {
	var signUpAccount internal.SignInAccount
	if err := json.NewDecoder(r.Body).Decode(&signUpAccount); err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_DECODE_DATA)
	}
	passwordByUsername, err := s.Storage.GetPasswordByUsername(signUpAccount.Username)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_SIGN_IN_USERNAME)
	}

	if util.DoPasswordsMatch(passwordByUsername, signUpAccount.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(6 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		})
		signedString, err := token.SignedString([]byte(signingKey))
		if err != nil {
			return response_error.New(err, http.StatusInternalServerError, ERR_SIGN_IN_PASSWORD)
		}
		return middleware.WriteJson(w, http.StatusOK, signedString)
	}
	return response_error.New(err, http.StatusBadRequest, ERR_HANDLE_TOKEN)
}

func (s *AccountApiServer) getAccounts(w http.ResponseWriter, r *http.Request) *response_error.Error {
	accounts, err := s.Storage.GetAllAccounts()
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_GET_ALL)
	}
	if len(accounts) == 0 {
		return middleware.WriteJson(w, http.StatusNoContent, MSG_NO_ACCOUNTS)
	}
	return middleware.WriteJson(w, http.StatusOK, accounts)
}

func (s *AccountApiServer) getAccountById(w http.ResponseWriter, r *http.Request) *response_error.Error {
	id := mux.Vars(r)["account_id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_CANNOT_DECODE_ID)
	}
	account, err := s.Storage.GetAccountById(idInt)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_CANNOT_GET_ACCOUNT)
	}
	emptyAccount := internal.Account{}
	if account == emptyAccount {
		return middleware.WriteJson(w, http.StatusNoContent, ERR_CANNOT_GET_ACCOUNT)
	}
	return middleware.WriteJson(w, http.StatusOK, account)
}

func (s *AccountApiServer) updateAccountById(w http.ResponseWriter, r *http.Request) *response_error.Error {
	id := mux.Vars(r)["account_id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_CANNOT_DECODE_ID)
	}
	var modAccount internal.SignUpAccount
	err = json.NewDecoder(r.Body).Decode(&modAccount)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_DECODE_DATA)
	}

	existingAccount, err := s.Storage.GetAccountById(idInt)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_CANNOT_GET_ACCOUNT_ID)
	}

	if util.DoPasswordsMatch(
		existingAccount.Password,
		modAccount.Password,
	) {
		modAccount.Password = existingAccount.Password
	} else {
		modPassword, err := util.HashPassword(modAccount.Password)
		if err != nil {
			return response_error.New(err, http.StatusInternalServerError, ERR_PASSWORD_MATCH)
		}
		modAccount.Password = modPassword
	}
	account := internal.AccountFrom(idInt, &modAccount)
	err = s.Storage.UpdateById(*account)
	if err != nil {
		return middleware.WriteJson(w, http.StatusNotFound, ERR_CANNOT_GET_ACCOUNT_ID)
	}
	return middleware.WriteJson(w, http.StatusOK, MSG_SUCCESS)
}

func (s *AccountApiServer) deleteAccountById(w http.ResponseWriter, r *http.Request) *response_error.Error {
	id := mux.Vars(r)["account_id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_CANNOT_DECODE_ID)
	}
	err = s.Storage.DeleteById(idInt)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_CANNOT_GET_ACCOUNT_ID)
	}
	return middleware.WriteJson(w, http.StatusOK, MSG_SUCCESS)
}

func (s *AccountApiServer) getAccountsByMask(w http.ResponseWriter, r *http.Request) *response_error.Error {
	var searchAccount internal.AccountSearch
	if err := json.NewDecoder(r.Body).Decode(&searchAccount); err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_DECODE_DATA)
	}
	accounts, err := s.Storage.GetAccountsByMask(searchAccount)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_GET_FIND_ACCOUNTS)
	}
	if len(accounts) == 0 {
		return middleware.WriteJson(w, http.StatusNoContent, MSG_NO_ACCOUNTS)
	}
	return middleware.WriteJson(w, http.StatusOK, accounts)
}
