package apigw

import (
	"encoding/json"
	"log"
	"net/http"
	"verbio/nlu"
)

func (api *API) getNLUClient() nlu.NLUService {
	return nlu.NewNLUService("nlu", api.getBaseClient())
}

func (api *API) rootNLU(w http.ResponseWriter, r *http.Request) error {
	//Decode request from json and forward to nlu
	nreq := &nlu.ProcessRequest{}
	if err := json.NewDecoder(r.Body).Decode(nreq); err != nil {
		log.Println("Cannot decode request")
		return err
	}
	resp, err := api.getNLUClient().Process(r.Context(), nreq)
	if err != nil {
		log.Printf("Error while connecting to nlu: %s", err)
		return err
	}
	return jsonResponse(w, resp)
}
