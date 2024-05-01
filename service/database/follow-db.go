package database

func (db *appdbimpl) InsertFollow() (string, error) {

	err := db.c.Exec(`
	INSERT OR IGNORE INTO Follow(IdUser,IdUserFollowed) 
	VALUES(?)`,
	)

	return name, err
}
