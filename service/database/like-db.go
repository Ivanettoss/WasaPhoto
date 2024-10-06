package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) InsertLike(idUserLike int, idPhoto int) error {

	_, err := db.c.Exec(`

	INSERT OR IGNORE INTO Like(IdUser,IdPhoto) 
		VALUES(?,?)`, idUserLike, idPhoto)

	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) DeleteLike(idUserPerforming int, idPhoto int) error {
	err := db.GetLike(idUserPerforming, idPhoto)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`
		DELETE FROM Like
		WHERE IdPhoto=? and IdUser=?`, idPhoto, idUserPerforming)

	if err != nil {
		return err
	}

	return nil

}

func (db *appdbimpl) GetLike(idUserPerforming int, idPhoto int) error {

	_, err := db.c.Query(`
	SELECT IdUser,IdPhoto
	FROM Like
	WHERE IdPhoto=? and IdUser=?`, idPhoto, idUserPerforming)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrLikeNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) CountLikes(idPhoto int) (int, error) {
	nlikes := 0
	err := db.c.QueryRow(`
	SELECT COUNT(*)
	FROM Like
	WHERE IdPhoto=?`, idPhoto).Scan(&nlikes)

	if errors.Is(err, sql.ErrNoRows) {
		return nlikes, ErrPhotoNotFound
	}

	return nlikes, err
}

func (db *appdbimpl) CheckLike(idUserPerforming int, idPhoto int) (bool, error) {
	var temp int
	err := db.c.QueryRow(`
	SELECT IdUser
	FROM Like
	WHERE IdPhoto=? and IdUser=?`, idPhoto, idUserPerforming).Scan(&temp)

	if err == sql.ErrNoRows {

		return false, nil

	} else if err != nil {

		return false, ErrLikeNotFound
	}

	// Se abbiamo trovato una riga, restituisci true
	return true, nil
}
