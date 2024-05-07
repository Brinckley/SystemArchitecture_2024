package middleware

import (
	"Gateway/internal/server/response_error"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
)

const (
	authHeader             = "Authorization"
	signingKey             = "wi2omp{wr2o*#3he3uJls@hsl3sa48"
	ERR_UNAUTHORIZED       = "unauthorized user"
	ERR_WRONG_TOKEN_FORMAT = "wrong token format"
)

type apiAuthorizedFunc func(w http.ResponseWriter, r *http.Request, accountId string) *response_error.Error

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
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

		if err := f(w, r, claims.UserId); err != nil {
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
