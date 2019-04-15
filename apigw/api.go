package apigw

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/micro/go-micro/metadata"
)

var ErrInvalid = errors.New("Invalid request")

type API struct {
	consulAddr string
}

func NewAPI(consulAddr string) *API {
	return &API{consulAddr}
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	head, _ := shiftPath(r.URL.Path)
	switch head {
	case "login":
		api.doAuthLoginCall(w, r)
	case "register":
		api.doAuthRegisterCall(w, r)
	case "":
		http.Error(w, "Invalid location", http.StatusBadRequest)
	default:
		err := api.validateAndServe(head, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func (api *API) validateAndServe(svc string, w http.ResponseWriter, r *http.Request) error {
	user, err := api.doAuthValidateCall(w, r)
	if err != nil {
		log.Println("Auth token is invalid")
		http.Error(w, "Invalid", http.StatusUnauthorized)
		return nil
	}
	md, ok := metadata.FromContext(r.Context())
	if !ok {
		md = metadata.Metadata{}
	}
	md["user"] = user.UserId
	ctx := metadata.NewContext(r.Context(), md)
	r = r.WithContext(ctx)
	switch svc {
	case "nlu":
		return api.rootNLU(w, r)
	case "dialog":
		return api.rootDM(w, r)
	}
	fmt.Println("Invalid service")
	return ErrInvalid
}
