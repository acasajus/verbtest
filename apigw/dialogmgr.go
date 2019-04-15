package apigw

import (
	"encoding/json"
	"log"
	"net/http"
	"verbio/dialogmgr"
)

func (api *API) getDMClient() dialogmgr.DialogMgrService {
	return dialogmgr.NewDialogMgrService("dialogmgr", api.getBaseClient())
}

func (api *API) rootDM(w http.ResponseWriter, r *http.Request) error {
	nreq := &dialogmgr.MessageRequest{}
	if err := json.NewDecoder(r.Body).Decode(nreq); err != nil {
		log.Println("Cannot decode request")
		return err
	}
	resp, err := api.getDMClient().Message(r.Context(), nreq)
	if err != nil {
		log.Printf("Error while connecting dialogmg: %s", err)
		return err
	}
	return jsonResponse(w, resp)
}
