package apigw

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
	client "github.com/micro/go-micro/client"
	consulreg "github.com/micro/go-micro/registry/consul"
)

func (api *API) getBaseClient() client.Client {
	conf := consulapi.Config{Address: api.consulAddr}
	reg := consulreg.NewRegistry(consulreg.Config(&conf))
	return client.NewClient(client.Registry(reg))
}

func jsonResponse(w http.ResponseWriter, obj interface{}) error {
	b := bytes.NewBuffer(nil)
	if err := json.NewEncoder(b).Encode(obj); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(b.Bytes())))
	w.WriteHeader(http.StatusOK)
	b.WriteTo(w)
	return nil
}
