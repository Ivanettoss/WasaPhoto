package database

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) BanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := UserAuthentication("u_name", w, r, ps)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uToBan := ps.ByName("user_to_ban")

	//get the user to follow from the username
	id, err := GetId(uToBan)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//insert it in db in ban table
	err := rt.db.InsertBan(components.Id, uToFollow)

}

func (rt *_router) UnBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//authenticate the user
	user, err := UserAuthentication("u_name", w, r, ps)
	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uToBan := ps.ByName("user_to_ban")

	//get the user to unban from the username
	id, err := GetId(uToBan)

	if err != nil {
		rt.baseLogger.WithError(err).Warning("Error reading Json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//insert it in db in ban table
	err := rt.db.DeleteBan(uToBan)

}
