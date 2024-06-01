package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)



func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the username of the profile owner
	profileOwner:=ps.ByName("u_name")

	token, err := GetBearerToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//get the profile owner id 
	profileOwnerId,err:=rt.db.GetId(profileOwner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get the username of the user performing action (me sa non me serve)
	/*
	userPerforming,err:=rt.db.GetUser(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}*/

	// check if the visitor is banned by the owner 
	err=rt.db.BanCheck(profileOwnerId,token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//lets build the profile stream
	var profile components.Profile

	// get the user photos 

	//get the followers number 
	profile.NFollowers,err=rt.db.GetFollowersNumber(profileOwnerId)

    // get the followed number 
	profile.NFollowed,err=rt.db.GetFollowedNumber(profileOwnerId)

	// get the post number (number of photo posted)
	profile.NPost,err=rt.db.GetPostedPhotoNumber(profileOwnerId)

	// get the follow state

	// get the ban state 
}