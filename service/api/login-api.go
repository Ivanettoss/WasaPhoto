package api

import (
	"net/http"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

//camel per le funzioni
func (rt *_router) login (w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext){
	var username string
	err:=json.NewDecoder(r.Body).Decode(&username)
	_=r.Body.Close

	if err!=nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	// call the function nel db che crea l'utente e lo inserta nel db 


}