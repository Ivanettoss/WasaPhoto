package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//user auth

	user, err := rt.UserAuthentication("u_name", w, r, ps)

	fmt.Println("utente post autenticazione", user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// rememeber to manage the user empty(ID 0 and username "" problem)
	var photo components.Photo

	// photo in base64 from the body saved in the variable
	err = json.NewDecoder(r.Body).Decode(&photo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update the struct variables
	photo.Username = user.Username

	photo.UploadDataTime = time.Now().Format("2006-01-02 15:04:05")

	//insert the photo into th db
	err = rt.db.InsertPhoto(user.Id, photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// return the new created photo
	fmt.Print("gestisci l'encode")
	_ = json.NewEncoder(w).Encode(photo)
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//user auth
	user, err := rt.UserAuthentication("u_name", w, r, ps)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photoId := ps.ByName("photo_id")

	//cast the id, i need an int
	photoIdInt, err := strconv.Atoi(photoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get the photo obejct to dfelete
	photoToDelete, err := rt.db.GetPhoto(photoIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if the pic is uploaded by the right user

	if user.Username != photoToDelete.Username {

		http.Error(w, database.ErrPageNotFound.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.DeletePhoto(photoIdInt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

}
