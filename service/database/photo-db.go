package database

import (
	"database/sql"
	"errors"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
)

func (db *appdbimpl) InsertPhoto(IdUser int, photo components.Photo) error {

	_, err := db.c.Exec(`

	INSERT OR IGNORE INTO Photo(IdUser,DateTime,Path) 
		VALUES(?,?,?)`, IdUser, photo.UploadDataTime, photo.PhotoBytes)

	if err != nil {
		return err
	}

	return nil

}

func (db *appdbimpl) DeletePhoto(photoId int) error {

	_, err := db.GetPhoto(photoId)

	if err != nil {
		return err
	}
	_, err = db.c.Exec(`
		DELETE FROM Photo
		WHERE IdPhoto=?`, photoId)

	if err != nil {
		return err
	}

	return nil

}

func (db *appdbimpl) GetPhoto(photoId int) (components.Photo, error) {

	var photo components.Photo
	err := db.c.QueryRow(`
	SELECT IdPhoto,Username,DateTime ,Path 
	FROM Photo,User
	WHERE Photo.IdPhoto=? and User.IdUser=Photo.IdUser`, photoId).Scan(&photo.IdPhoto, &photo.Username, &photo.UploadDataTime, &photo.PhotoBytes)

	if errors.Is(err, sql.ErrNoRows) {
		return photo, ErrPhotoNotFound
	}

	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (db *appdbimpl) GetPhotoOwnerId(photoId int) (int, error) {

	var idOwner int

	err := db.c.QueryRow(`
	SELECT IdUser
	FROM Photo
	WHERE Photo.IdPhoto=? `, photoId).Scan(&idOwner)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, ErrPhotoNotFound
	}

	if err != nil {
		return 0, err
	}
	return idOwner, nil
}

func (db *appdbimpl) GetPostedPhotoNumber(idUser int) (int, error) {
	var postNumber int
	err := db.c.QueryRow(`
	SELECT COUNT(IdPhoto)
	FROM Photo
	WHERE Photo.IdUser=?`, idUser).Scan(&postNumber)

	if err != nil {
		return postNumber, err
	}
	return postNumber, nil
}

func (db *appdbimpl) GetUserPhotos(idUserPerforming int, username string) ([]components.Photo, error) {
	var photoList []components.Photo

	photoRows, err := db.c.Query(`
	SELECT Photo.IdPhoto,User.Username, Photo.DateTime,Photo.Path
	FROM Photo,User
	WHERE Photo.IdUser=User.IdUser and Username=?`, username)

	if err != nil {
		return photoList, err
	}

	defer photoRows.Close()

	for photoRows.Next() {
		var photo components.Photo
		err = photoRows.Scan(&photo.IdPhoto, &photo.Username, &photo.UploadDataTime, &photo.PhotoBytes)
		if err != nil {
			return photoList, err
		}
		photo.NLikes, err = db.CountLikes(photo.IdPhoto)
		if err != nil {
			return photoList, ErrPhotoNotFound
		}

		photo.IsLiked, err = db.CheckLike(idUserPerforming, photo.IdPhoto)
		if err != nil {
			return photoList, err
		}

		photo.Comments, err = db.GetCommentList(photo.IdPhoto)
		if err != nil {
			return photoList, err
		}

		photo.NComments = len(photo.Comments)

		photoList = append(photoList, photo)
	}

	return photoList, nil
}
