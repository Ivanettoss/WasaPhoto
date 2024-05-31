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

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	//get the comment object from the request body
	var comment components.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(comment)

	//get the comment username from the body
	userNameFromComment := comment.User.Username
	fmt.Println("username dal body", userNameFromComment)

	// get the id from the db linked to the username
	idFromDatabase, err := rt.db.GetId(userNameFromComment)
	fmt.Println("id dell'username dal db", idFromDatabase)

	fmt.Println("id dell'username nel body")
	//check if the ids match
	if idFromDatabase != comment.User.Id {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusInternalServerError)
		return
	}

	// now we check if the user that want to comment is authorized
	err = CheckAuth(comment.User, r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//get the photo id from the comment
	photoId, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("id della foto dal body", photoId)

	//id of the photo owner
	userOwnerId, err := rt.db.GetId(ps.ByName("u_name"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("username proprietario della foto dal path", userOwnerId)

	// get the id of the owner from a photo
	photoOwnerId, err := rt.db.GetPhotoOwnerId(photoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("username proprietario della foto dal db  ", userOwnerId)

	//check if the user is the real owner
	if photoOwnerId != userOwnerId {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusInternalServerError)
		return
	}

	// now we can insert the comment into the database

	// first set the correct datatime value
	comment.UploadDataTime = time.Now().Format("2006-01-02 15:04:05")
	comment.IdPhoto = photoId

	err = rt.db.InsertComment(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// return the newly created comment
	_ = json.NewEncoder(w).Encode(comment)
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get the id from the path
	commentId, err := strconv.Atoi(ps.ByName("comment_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("id del commento dal path", commentId)

	// get the comment object from db
	commentToDelete, err := rt.db.GetComment(commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(" commento selezionato", commentToDelete)
	// we want to see if the user that is performing the action is the comment author

	// get token
	token, err := GetBearerToken(r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if commentToDelete.User.Id != token {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusInternalServerError)
		return
	}

	// check if the photo is consistent

	// get the user owner of the photo from the photo
	photoId, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userPhotoOwnerId, err := rt.db.GetPhotoOwnerId(photoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//get the user owner of the photo from the path
	username := ps.ByName("u_name")
	userId, err := rt.db.GetId(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the ids are the same
	if userId != userPhotoOwnerId {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusInternalServerError)
		return
	}

	// now we can delete the comment
	err = rt.db.DeleteComment(commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	// return the removed comment
	_ = json.NewEncoder(w).Encode(commentToDelete)

}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get the bearer token
	token, err := GetBearerToken(r.Header.Get("Authorization"))

	fmt.Println("token", token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the user that is performing the action
	userPerforming, err := rt.db.GetUser(token)
	fmt.Print("user performing", userPerforming)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("user", userPerforming)
	//get the photo owner from path
	photoOwner := ps.ByName("u_name")
	photoOwnerId, err := rt.db.GetId(photoOwner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("photo owner", photoOwnerId)
	//check if the user is banned
	err = rt.db.BanCheck(photoOwnerId, userPerforming.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// now check if the photo is connsistent
	fmt.Println("qui parte il check della consistenza")
	//get the photo id from the path
	photoId, err := strconv.Atoi(ps.ByName("photo_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//get the user owner id from the photo id
	ownerIdFromPhoto, err := rt.db.GetPhotoOwnerId(photoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("proprietario della foto dal path", photoOwnerId)
	fmt.Println("proprietario della foto dall'object photo", ownerIdFromPhoto)

	if photoOwnerId != ownerIdFromPhoto {
		http.Error(w, database.ErrPageNotFound.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("fino a qui tutto bene")
	commentList, err := rt.db.GetCommentList(photoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
	// return the new object l
	_ = json.NewEncoder(w).Encode(commentList)

}
