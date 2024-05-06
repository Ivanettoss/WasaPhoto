package database

func (db *appdbimpl) InsertBan(idUserPerforming int, idUserToBan int) error {

	_, err := db.c.Exec(`
	INSERT OR IGNORE INTO Follow(IdUser,IdUserBanned) 
	VALUES(?,?)`, idUserPerforming, idUserToBan)

	return err
}

func (db *appdbimpl) DeleteBan(idUserToUnBan int) error {

	_, err := db.c.Exec(`
	DELETE FROM Ban
	WHERE IdUserBanned=?`, idUserToUnBan)

	return err
}
