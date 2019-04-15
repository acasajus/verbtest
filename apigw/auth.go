package apigw

import (
	"log"
	"net/http"
	"strings"
	"verbio/auth"
)

func (api *API) getAuthClient() auth.AuthService {
	return auth.NewAuthService("auth", api.getBaseClient())
}

func (api *API) doAuthLoginCall(w http.ResponseWriter, r *http.Request) error {
	//TODO: Parse and forward to auth svc
	return nil
}

func (api *API) doAuthRegisterCall(w http.ResponseWriter, r *http.Request) error {
	//TODO: Parse and forward to auth svc
	return nil
}

func (api *API) doAuthValidateCall(w http.ResponseWriter, r *http.Request) (*auth.ValidateResponse, error) {
	//Send validation request with the token extracted from the Authorization: Bearer XXXX
	authHdr := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authHdr) < 2 || authHdr[0] != "Bearer" {
		log.Printf("Invalid auth header is %v", authHdr)
		return nil, auth.ErrUnauthorized
	}
	log.Printf("Valid auth header is %v", authHdr)
	req := auth.ValidateRequest{Token: authHdr[1]}
	svc := api.getAuthClient()
	return svc.Validate(r.Context(), &req)
}
