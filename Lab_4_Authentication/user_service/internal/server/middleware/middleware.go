package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"user_service/internal/server/response_error"
)

const (
	authHeader             = "Authorization"
	signingKey             = "wi2omp{wr2o*#3he3uJls@hsl3sa48"
	ERR_UNAUTHORIZED       = "unauthorized user"
	ERR_WRONG_TOKEN_FORMAT = "wrong token format"
)

func WriteJsonFromResponse(w http.ResponseWriter, status int, r *http.Response) *response_error.Error {
	w.WriteHeader(status)
	_, err := io.Copy(w, r.Body)
	if err != nil {
		return response_error.New(err, http.StatusInternalServerError, "cannot copy body")
	}
	return nil
}

func WriteJson(w http.ResponseWriter, status int, content any) *response_error.Error {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		log.Println(err)
		return response_error.New(err, http.StatusInternalServerError, "cannot encode error")
	}
	return nil
}

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("account service received request method: %s, header: %s, body: %s\n", r.Method, r.Header, r.Body)
		next.ServeHTTP(w, r)
	})
}

type apiFunc func(w http.ResponseWriter, r *http.Request) *response_error.Error

func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := WriteJson(w, err.StatusCode(), err.Unwrap())
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}

type apiAuthorizedFunc func(w http.ResponseWriter, r *http.Request, accountId string) *response_error.Error

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func MakeHTTPAuthedHandleFunc(f apiAuthorizedFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := getValidatedToken(w, r)
		if err != nil {
			err = WriteJson(w, err.StatusCode(), err.Unwrap())
		}

		tokenWithClaims, errJwt := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(signingKey), nil
		})
		if errJwt != nil {
			err = WriteJson(w, http.StatusUnauthorized, ERR_WRONG_TOKEN_FORMAT)
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
			}
			return
		}

		claims, ok := tokenWithClaims.Claims.(*tokenClaims)
		if !ok {
			err = WriteJson(w, http.StatusUnauthorized, ERR_WRONG_TOKEN_FORMAT)
		}

		if err := f(w, r, strconv.Itoa(claims.UserId)); err != nil {
			err := WriteJson(w, err.StatusCode(), err.Unwrap())
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}

func getValidatedToken(w http.ResponseWriter, req *http.Request) (string, *response_error.Error) {
	header := req.Header.Get(authHeader)
	if header == "" {
		err := WriteJson(w, http.StatusUnauthorized, ERR_UNAUTHORIZED)
		if err != nil {
			log.Println(fmt.Errorf("unable to write error data error : %s", err))
			return "", err
		}
		return "", nil
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err := WriteJson(w, http.StatusUnauthorized, ERR_WRONG_TOKEN_FORMAT)
		if err != nil {
			log.Println(fmt.Errorf("unable to write error data error : %s", err))
			return "", err
		}
		return "", nil
	}
	return headerParts[1], nil
}
