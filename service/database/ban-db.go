package database

func (db *appdbimpl) InsertBan(idUserPerforming int, idUserToBan int) error {

	_, err := db.c.Exec(`
	INSERT OR IGNORE INTO Ban(IdUser,IdUserBanned) 
	VALUES(?,?)`, idUserPerforming, idUserToBan)

	return err
}

func (db *appdbimpl) DeleteBan(idUserPerforming int, idUserToUnBan int) error {

	_, err := db.c.Exec(`
	DELETE FROM Ban
	WHERE IdUserBanned=? and IdUser=?`, idUserToUnBan, idUserPerforming)

	return err
}

func (db *appdbimpl) GetBannedList(idUserPerforming int) ([]string, error) {
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
