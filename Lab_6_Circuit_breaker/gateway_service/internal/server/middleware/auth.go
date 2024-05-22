package middleware

import (
	"Gateway/internal/server/response_error"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

const (
	authHeader             = "Authorization"
	signingKey             = "wi2omp{wr2o*#3he3uJls@hsl3sa48"
	ERR_UNAUTHORIZED       = "unauthorized user"
	ERR_WRONG_TOKEN_FORMAT = "wrong token format"
	TOKEN_HEADER_NAME      = "Authorization"
	HEADER_USER_ID         = "User-Id"
)

type apiAuthorizedFunc func(w http.ResponseWriter, r *http.Request, accountId string) *response_error.Error

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func MakeHTTPAuthedHandleFunc(f apiAuthorizedFunc, accountUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(TOKEN_HEADER_NAME)
		request, err := http.NewRequest(http.MethodGet, accountUrl+"/auth", nil)
		if err != nil {
			err := WriteJson(w, http.StatusBadRequest, fmt.Sprintf("cannot create auth request %s", err))
			log.Println("cannot create request ", err)
			return
		}
		request.Header.Set("Authorization", token)
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			err := WriteJson(w, http.StatusBadRequest, "cannot exec auth request %s")
			log.Println("cannot create request ", err)
			return
		}
		log.Println("answer from account : ", resp)

		if resp.StatusCode != http.StatusOK {
			err := WriteJson(w, resp.StatusCode, resp.Body)
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
			}
			log.Println("cannot send request for id to the account_service request : ", resp)
			return
		}

		userId := resp.Header.Get(HEADER_USER_ID)
		if err := f(w, r, userId); err != nil {
			err := WriteJson(w, err.StatusCode(), err.Unwrap())
			if err != nil {
				log.Println(fmt.Errorf("unable to write error data error : %s", err))
				return
			}
		}
	}
}
