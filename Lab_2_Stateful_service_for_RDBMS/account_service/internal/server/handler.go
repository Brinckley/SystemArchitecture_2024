package server

import (
	"account_service/internal"
	"account_service/internal/util"
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
	if len(accounts) == 0 {
		return writeJson(w, http.StatusNoContent, []internal.Account{})
	}
	return writeJson(w, http.StatusOK, accounts)
}

func (s *AccountApiServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	var createAccountReq internal.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil {
		return writeJson(w, http.StatusBadRequest, "cannot decode account data")
	}
	newPassword, err := util.HashPassword(createAccountReq.Password)
	if err != nil {
		return err
	}
	createAccountReq.Password = newPassword
	log.Println(createAccountReq)
	accountId, err := s.Storage.CreateAccount(&createAccountReq)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, "cannot create account or already exists")
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
		log.Println(err)
		return writeJson(w, http.StatusNotFound, "cannot find account")
	}
	if account == nil {
		return writeJson(w, http.StatusNoContent, internal.Account{})
	}
	return writeJson(w, http.StatusOK, account)
}

func (s *AccountApiServer) updateAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusBadRequest, "cannot decode input data")
	}
	var modAccount internal.CreateAccountRequest
	err = json.NewDecoder(r.Body).Decode(&modAccount)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusBadRequest, "cannot decode input data")
	}

	existingAccount, err := s.Storage.GetAccountById(idInt)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusNotFound, "cannot find the account or update it")
	}

	if util.DoPasswordsMatch(
		existingAccount.Password,
		modAccount.Password,
	) {
		log.Println("Password match found in update")
		modAccount.Password = existingAccount.Password
	} else {
		modPassword, err := util.HashPassword(modAccount.Password)
		if err != nil {
			return err
		}
		modAccount.Password = modPassword
	}
	account := internal.AccountFrom(idInt, &modAccount)
	accountModified, err := s.Storage.UpdateAccount(account)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusNotFound, "cannot find the account or update it")
	}
	return writeJson(w, http.StatusOK, accountModified)
}

func (s *AccountApiServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusBadRequest, "cannot decode input data")
	}
	deletedId, err := s.Storage.DeleteAccount(idInt)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusNotFound, "cannot find the account or delete it")
	}
	return writeJson(w, http.StatusOK, deletedId)
}

func (s *AccountApiServer) getAccountsByMask(w http.ResponseWriter, r *http.Request) error {
	log.Println("MASK")
	var searchAccount internal.AccountSearch
	if err := json.NewDecoder(r.Body).Decode(&searchAccount); err != nil {
		log.Println(err)
		return writeJson(w, http.StatusBadRequest, "cannot decode account data")
	}
	accounts, err := s.Storage.GetAccountsByMask(&searchAccount)
	if err != nil {
		return err
	}
	log.Println(accounts)
	if len(accounts) == 0 {
		return writeJson(w, http.StatusNoContent, []internal.Account{})
	}
	return writeJson(w, http.StatusOK, accounts)
}
