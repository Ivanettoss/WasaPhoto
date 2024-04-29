package database

func (db *appdbimpl) login ( username string) error {
	_, err := db.c.Exec("INSERT OR IGNORE INTO User(Username) VALUES(\'?\')", username)
	return err
}