package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// camel case per le funzioni
func (rt *_router) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var username string
	var id int

	err := json.NewDecoder(r.Body).Decode(&username)
	_ = r.Body.Close

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// call the function nel db che crea l'utente e lo inserta nel db
	// ho il suo id
	id, err = rt.db.InsertUser(username)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("ERRORE INSERIMENTO NEL DB ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user components.User

	// voglio la struct del determinato utente
	user, err = rt.db.GetUser(id)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("// da fare  ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201

	// return the new object l
	_ = json.NewEncoder(w).Encode(user)

}
