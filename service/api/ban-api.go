package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := rt.UserAuthentication("u_name", w, r, ps)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uToBan := ps.ByName("user_to_ban")

	//get the user's Id to ban from the username
	id, err := rt.db.GetId(uToBan)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//insert it in db in ban table
	err = rt.db.InsertBan(user.Id, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204
}

func (rt *_router) unBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := rt.UserAuthentication("u_name", w, r, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uToBan := ps.ByName("user_to_ban")

	//get the user to unban from the username
	id, err := rt.db.GetId(uToBan)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//insert it in db in ban table
	err = rt.db.DeleteBan(user.Id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent) // 204

}

func (rt *_router) banList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//authenticate the user
	user, err := rt.UserAuthentication("u_name", w, r, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	banList, err := rt.db.GetBannedList(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// return the new object l
	_ = json.NewEncoder(w).Encode(banList)

}
