package database

func (db *appdbimpl) InsertUser(username string) (error,int) {
	var id int 
 
	// if not exist insert the user into the db 
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO User(Username) 
		VALUES(?)`, username)

	if err != nil {
		return err,0
	}

	// select the id from the last one inserted 
	err = db.c.QueryRow(`
		SELECT Id
		FROM User
		WHERE User.username=?`,username).Scan(&id)

	if err != nil {
			return err,0
	}

	// return the id from the last inserted 
	return nil,id
}



