package server

import (
	"account_service/internal"
	"account_service/internal/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (s *AccountApiServer) getAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.Storage.GetAll(*s.Ctx)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, []internal.Account{})
	}
	if len(accounts) == 0 {
		return writeJson(w, http.StatusNoContent, []internal.Account{})
	}
	return writeJson(w, http.StatusOK, accounts)
}

func (s *AccountApiServer) createAccount(w http.ResponseWriter, r *http.Request) error {
	var accountDto internal.AccountDto
	if err := json.NewDecoder(r.Body).Decode(&accountDto); err != nil {
		return writeJson(w, http.StatusBadRequest, err)
	}

	newPassword, err := util.HashPassword(accountDto.Password)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, fmt.Sprintf("cannot hash password error %v", err))
	}
	accountDto.Password = newPassword

	accountId, err := s.Storage.Create(*s.Ctx, accountDto)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, err)
	}
	return writeJson(w, http.StatusOK, accountId)
}

func (s *AccountApiServer) getAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	account, err := s.Storage.GetById(*s.Ctx, id)
	if err != nil {
		log.Println(err)
		return writeJson(w, http.StatusNoContent, "cannot find account")
	}
	return writeJson(w, http.StatusOK, account)
}

func (s *AccountApiServer) updateAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	var modAccount internal.Account
	err := json.NewDecoder(r.Body).Decode(&modAccount)
	if err != nil {
		return writeJson(w, http.StatusBadRequest, fmt.Sprintf("cannot decode input data error %v", err))
	}
	modAccount.Id = id

	err = s.Storage.Update(*s.Ctx, modAccount)
	if err != nil {
		return writeJson(w, http.StatusNoContent, fmt.Sprintf("cannot find the account or update it error %v", err))
	}
	return writeJson(w, http.StatusOK, "account modified")
}

func (s *AccountApiServer) deleteAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := s.Storage.Delete(*s.Ctx, id)
	if err != nil {
		return writeJson(w, http.StatusNoContent, fmt.Sprintf("cannot find the account or delete it error %v", err))
	}
	return writeJson(w, http.StatusOK, "account deleted")
}

func (s *AccountApiServer) getAccountsByMask(w http.ResponseWriter, r *http.Request) error {
	var searchAccount internal.AccountSearch
	if err := json.NewDecoder(r.Body).Decode(&searchAccount); err != nil {
		return writeJson(w, http.StatusBadRequest, fmt.Sprintf("cannot decode account data error %v", err))
	}

	accounts, err := s.Storage.GetByMask(*s.Ctx, searchAccount)
	if err != nil {
		return writeJson(w, http.StatusNoContent, fmt.Sprintf("cannot get accounts by mask error %v", err))
	}
	return writeJson(w, http.StatusOK, accounts)
}
