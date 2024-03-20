package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"user_service/internal"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(fmt.Errorf("unable to write data error : %s", err))
			err := writeJson(w, http.StatusBadRequest, "Unable to handle data")
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}

func (s *ApiServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	var createAccountReq internal.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return err
	}
	log.Println(createAccountReq)
	accountId, err := s.Storage.CreateAccount(&createAccountReq)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, accountId)
}

func (s *ApiServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	account, err := s.Storage.GetAccountById(idInt)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, account)
}

func (s *ApiServer) modifyAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	var modAccount internal.CreateAccountRequest
	err = json.NewDecoder(r.Body).Decode(&modAccount)
	if err != nil {
		return err
	}
	account := internal.AccountFrom(idInt, &modAccount)
	accountModified, err := s.Storage.UpdateAccount(account)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, accountModified)
}

func (s *ApiServer) getAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.Storage.GetAccounts()
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, accounts)
}

func (s *ApiServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = s.Storage.DeleteAccount(idInt)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, idInt)
}
