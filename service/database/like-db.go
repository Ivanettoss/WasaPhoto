package database

import (
	"database/sql"
	"errors"
	
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
