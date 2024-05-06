package database

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) FollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := UserAuthentication("u_name", w, r, ps)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uToFollow := ps.ByName("name_to_follow")

	//get the user to follow from the username
	id, err := GetId(uToFollow)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//insert it in db in follow table
	err := rt.db.InsertFollow(components.Id, uToFollow)

}

func (rt *_router) UnFollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := UserAuthentication("u_name", w, r, ps)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uToFollow := ps.ByName("name_to_follow")

	//get the user to follow from the username
	id, err := GetId(uToFollow)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//insert it in db in follow table
	err := rt.dd.DeleteFollow(uToFollow)

}

func (rt *_router) FollowersList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := UserAuthentication("u_name", w, r, ps)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.dd.getFollowersList(user)
	// da finire

}

func (rt *_router) FollowedList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := UserAuthentication("u_name", w, r, ps)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.dd.getFollowedList(user)

	/// da finire
}
