package database

func (db *appdbimpl) InsertUser(username string) (int, error) {
	var id int

	// if not exist insert the user into the db
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO User(Username) 
		VALUES(?)`, username)

	if err != nil {
		return 0, err
	}

	// select the id from the last one inserted
	err = db.c.QueryRow(`
		SELECT IdUser
		FROM User
		WHERE User.username=?`, username).Scan(&id)

	if err != nil {
		return 0, err
	}

	// return the id from the last inserted
	return id, nil
}
