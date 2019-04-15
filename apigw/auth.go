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
	authHdr := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authHdr) < 2 || authHdr[0] != "Bearer" {
		return nil, auth.ErrUnauthorized
	}
	log.Printf("Auth header is %s", authHdr[1])
	req := auth.ValidateRequest{Token: authHdr[1]}
	svc := api.getAuthClient()
	return svc.Validate(r.Context(), &req)
}
