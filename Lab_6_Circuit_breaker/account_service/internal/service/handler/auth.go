package handler

import (
	"account_service/internal"
	"account_service/internal/service/middleware"
	"account_service/internal/service/response_error"
	"account_service/internal/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ERR_UNAUTHORIZED        = "unauthorized user"
	ERR_WRONG_TOKEN_FORMAT  = "wrong token format"
	ERR_WRONG_HASH_FUNCTION = "wrong hash function"
	ERR_EMPTY_TOKEN         = "empty token"
	ERR_RETURNING_ID        = "cannot return id"
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

type tokenCustomClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (s *AccountApiServer) signInAccount(w http.ResponseWriter, r *http.Request) *response_error.Error {
	var signUpAccount internal.SignInAccount
	if err := json.NewDecoder(r.Body).Decode(&signUpAccount); err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_DECODE_DATA)
	}
	userId, passwordByUsername, err := s.Storage.GetPasswordByUsername(signUpAccount.Username)
	if err != nil {
		return response_error.New(err, http.StatusBadRequest, ERR_SIGN_IN_USERNAME)
	}

	if util.DoPasswordsMatch(passwordByUsername, signUpAccount.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenCustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(6 * time.Hour).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			UserId: strconv.Itoa(userId),
		})

		signedString, err := token.SignedString([]byte(signingKey))
		if err != nil {
			return response_error.New(err, http.StatusInternalServerError, ERR_SIGN_IN_PASSWORD)
		}
		w.Header().Add("Auth-token", signedString)
		return middleware.WriteWithTokenHeader(w, http.StatusOK)
	}
	return response_error.New(err, http.StatusBadRequest, ERR_HANDLE_TOKEN)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func (s *AccountApiServer) getAccountIdFromToken(w http.ResponseWriter, r *http.Request) *response_error.Error {
	token, err := getValidatedToken(r)
	if err != nil {
		return err
	}

	tokenWithClaims, errJwt := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(signingKey), nil
	})
	if errJwt != nil {
		return response_error.New(errJwt, http.StatusBadRequest, ERR_WRONG_HASH_FUNCTION)
	}

	claims, ok := tokenWithClaims.Claims.(*tokenClaims)
	if !ok {
		return response_error.New(err, http.StatusUnauthorized, ERR_WRONG_TOKEN_FORMAT)
	}
	w.Header().Add("User-Id", claims.UserId)
	return middleware.WriteWithIdHeader(w, http.StatusOK)
}

func getValidatedToken(req *http.Request) (string, *response_error.Error) {
	header := req.Header.Get("Authorization")
	if header == "" {
		return "", response_error.New(fmt.Errorf("empty authorization header"), http.StatusBadRequest, ERR_EMPTY_TOKEN)
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", response_error.New(fmt.Errorf("empty authorization header"), http.StatusBadRequest, ERR_WRONG_TOKEN_FORMAT)
	}
	return headerParts[1], nil
}
