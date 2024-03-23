package server

import (
	"account_service/internal"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (s *AccountApiServer) getAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.Storage.GetAccounts()
	if err != nil {
		return err
	}
	log.Println(accounts)
	return writeJson(w, http.StatusOK, accounts)
}

func (s *AccountApiServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	var createAccountReq internal.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return err
	}
	log.Println(createAccountReq)
	err := createAccountReq.SetPassword(createAccountReq.Password)
	if err != nil {
		return err
	}
	accountId, err := s.Storage.CreateAccount(&createAccountReq)
	if err != nil {
		return err
	}
	return writeJson(w, http.StatusOK, accountId)
}

func (s *AccountApiServer) getAccount(w http.ResponseWriter, r *http.Request) error {
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

func (s *AccountApiServer) updateAccount(w http.ResponseWriter, r *http.Request) error {
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
	err = modAccount.SetPassword(modAccount.Password)
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

func (s *AccountApiServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
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
