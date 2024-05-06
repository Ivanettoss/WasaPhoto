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

func (db *appdbimpl) getBannedList(idUserPerforming int) ([]string, error) {
	var BannedList []string

	BannedRows, err := db.c.Query(`
	SELECT Username
	FROM Ban, User
	WHERE Ban.IdUserBanned=User.IdUser and  Ban.IdUser=? `, idUserPerforming)

	defer BannedRows.Close()

	for BannedRows.Next() {
		var name string
		BannedRows.Scan(&name)
		BannedList = append(BannedList, name)
	}

	// format the list
	return BannedList, err

}
