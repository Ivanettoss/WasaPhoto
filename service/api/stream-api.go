package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// autenticate the user performing the action 
	user, err := rt.UserAuthentication("u_name", w, r, ps)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println("user id",user.Id)
	fmt.Println("user name",user.Username)
	currentStream, err := rt.db.GetStream(user.Id, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// return the new created stream
	_ = json.NewEncoder(w).Encode(currentStream)

}
