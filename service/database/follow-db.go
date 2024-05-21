package database

func (db *appdbimpl) InsertFollow(idUserPerforming int, idUserToFollow int) error {

	_, err := db.c.Exec(`
	INSERT OR IGNORE INTO Follow(IdUser,IdUserFollowed) 
	VALUES(?,?)`, idUserPerforming, idUserToFollow)

	return err
}

func (db *appdbimpl) DeleteFollow(idUserPerforming int, idUserToUnFollow int) error {

	_, err := db.c.Exec(`
	DELETE 
	FROM Follow
	WHERE IdUserFollowed=? and IdUser=?`, idUserToUnFollow, idUserPerforming)
	return err
}

func (db *appdbimpl) GetFollowersList(idUser int) ([]string, error) {
	var FollowersList []string

	FollowerRows, err := db.c.Query(`
	SELECT Username
	FROM Follow, User
	WHERE Follow.IdUser=User.IdUser and  Follow.IdUserFollowed=? `, idUser)

	defer FollowerRows.Close()

	for FollowerRows.Next() {
		var name string
		FollowerRows.Scan(&name)
		FollowersList = append(FollowersList, name)
	}

	// format the list
	return FollowersList, err

}

func (db *appdbimpl) GetFollowedList(idUser int) ([]string, error) {
	var FollowersList []string

	FollowerRows, err := db.c.Query(`
	SELECT Username
	FROM Follow, User
	WHERE Follow.IdUserFollowed=User.IdUser and  Follow.IdUser=? `, idUser)

	defer FollowerRows.Close()

	for FollowerRows.Next() {
		var name string
		FollowerRows.Scan(&name)
		FollowersList = append(FollowersList, name)
	}

	// format the list
	return FollowersList, err
}
