package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the username of the profile owner
	profileOwner := ps.ByName("u_name")

	// get the id of the user performing
	userPerformingId, err := GetBearerToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// get the profile owner id
	profileOwnerId, err := rt.db.GetId(profileOwner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bancheck, err := rt.db.BanCheck(profileOwnerId, userPerformingId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if bancheck {
		http.Error(w, "impossible to retrieve the resource", http.StatusInternalServerError)
		return
	}

	// lets build the profile stream
	var profile components.Profile

	profile.Username = profileOwner

	// get the user photos
	profile.Photos, err = rt.db.GetUserPhotos(userPerformingId, profileOwner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the followers number
	profile.NFollowers, err = rt.db.GetFollowersNumber(profileOwnerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the followed number
	profile.NFollowed, err = rt.db.GetFollowedNumber(profileOwnerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the post number (number of photo posted)
	profile.NPost, err = rt.db.GetPostedPhotoNumber(profileOwnerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the follow state
	profile.FollowState, err = rt.db.GetFollowState(userPerformingId, profileOwnerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the ban state
	profile.BanState, err = rt.db.BanCheck(userPerformingId, profileOwnerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the user profile
	_ = json.NewEncoder(w).Encode(profile)

}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var new_username components.Username

	// authenticate the user performing
	user, err := rt.UserAuthentication("u_name", w, r, ps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//  decode the new username from the body
	err = json.NewDecoder(r.Body).Decode(&new_username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//  update the username
	err = rt.db.SetUsername(user.Username, new_username.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	_ = json.NewEncoder(w).Encode(new_username)

}

func (rt *_router) getUsersList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	myName := ps.ByName("u_name")
	var usersFound components.Ulist
	userToFind := ps.ByName("user_to_find")
	usersFound, err := rt.db.SearchUsername(myName, userToFind)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(usersFound)

}
