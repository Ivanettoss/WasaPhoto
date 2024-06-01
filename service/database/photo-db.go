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
