package api

import (
	"encoding/json"

	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// authenticate the user

	user, err := rt.UserAuthentication("u_name", w, r, ps)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uToFollow := ps.ByName("name_to_follow")

	// get the user to follow from the username
	id, err := rt.db.GetId(uToFollow)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert it in db in follow table
	err = rt.db.InsertFollow(user.Id, id)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("to do")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the username of the user followed
	_ = json.NewEncoder(w).Encode(uToFollow)

}

func (rt *_router) unFollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// authenticate the user
	user, err := rt.UserAuthentication("u_name", w, r, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uToFollow := ps.ByName("name_to_follow")

	// get the user's Id to follow from the username
	id, err := rt.db.GetId(uToFollow)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// REMOVE it from follow table
	err = rt.db.DeleteFollow(user.Id, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204

}

