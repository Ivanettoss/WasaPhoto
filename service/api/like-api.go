package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// authenticate the user performing
	user, err := rt.UserAuthentication("like_name", w, r, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// id of the photo to like converted in int
	idPhotoToLike, err := strconv.Atoi(ps.ByName("photo_id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// id of the photo owner
	userOwnerId, err := rt.db.GetId(ps.ByName("u_name"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photoToLike, err := rt.db.GetPhoto(idPhotoToLike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get the id of the owner from a photo
	photoOwnerId, err := rt.db.GetPhotoOwnerId(idPhotoToLike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the user is the real owner
	if photoOwnerId != userOwnerId {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusInternalServerError)
		return
	}

	// insert the like
	err = rt.db.InsertLike(user.Id, idPhotoToLike)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update the like counter
	photoToLike.NLikes, err = rt.db.CountLikes(idPhotoToLike)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// authenticate the user performing
	user, err := rt.UserAuthentication("like_name", w, r, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// id of the photo to like converted in int
	idPhotoToLike, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// id of the photo owner
	userOwnerId, err := rt.db.GetId(ps.ByName("u_name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// photo object of the photo to like
	photoToLike, err := rt.db.GetPhoto(idPhotoToLike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get the id of the owner from a photo
	photoOwnerId, err := rt.db.GetPhotoOwnerId(idPhotoToLike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the user is the real owner
	if photoOwnerId != userOwnerId {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusBadRequest)
		return
	}

	// delete the like
	err = rt.db.DeleteLike(user.Id, idPhotoToLike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update the like counter
	photoToLike.NLikes, err = rt.db.CountLikes(idPhotoToLike)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
