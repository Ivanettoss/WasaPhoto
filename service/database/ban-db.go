package database

func (db *appdbimpl) InsertBan(idUserPerforming int, idUserToBan int) error {

	_, err := db.c.Exec(`
	INSERT OR IGNORE INTO Ban(IdUser,IdUserBanned) 
	VALUES(?,?)`, idUserPerforming, idUserToBan)

	if err != nil {
		return err
	}

	return err
}

func (db *appdbimpl) DeleteBan(idUserPerforming int, idUserToUnBan int) error {

	_, err := db.c.Exec(`
	DELETE FROM Ban
	WHERE IdUserBanned=? and IdUser=?`, idUserToUnBan, idUserPerforming)

	if err != nil {
		return err
	}

	return err
}

func (db *appdbimpl) GetBannedList(idUserPerforming int) ([]string, error) {
	var BannedList []string

	BannedRows, err := db.c.Query(`
	SELECT Username
	FROM Ban, User
	WHERE Ban.IdUserBanned=User.IdUser and  Ban.IdUser=? `, idUserPerforming)

	if err != nil {
		return BannedList, err
	}

	defer BannedRows.Close()

	for BannedRows.Next() {
		var name string
		err = BannedRows.Scan(&name)
		if err != nil {
			return BannedList, err
		}
		BannedList = append(BannedList, name)
	}

	// format the list
	return BannedList, err

}

func (db *appdbimpl) BanCheck(idOwner int, idUserPerforming int) (bool, error) {
	var bancheck bool

	err := db.c.QueryRow(`
	SELECT EXISTS(
		SELECT 1
		FROM Ban
		WHERE IdUser=?
		AND IdUserBanned=?
	)
	`, idOwner, idUserPerforming).Scan(&bancheck)

	if err != nil {

		return bancheck, err
	}

	return bancheck, nil
}
