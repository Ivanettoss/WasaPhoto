package database

func (db *appdbimpl) InsertFollow(idUserPerforming int, idUserToFollow int) error {

	_, err := db.c.Exec(`
	INSERT OR IGNORE INTO Follow(IdUser,IdUserFollowed) 
	VALUES(?,?)`, idUserPerforming, idUserToFollow)

	if err != nil {
		return err
	}
	return err
}

func (db *appdbimpl) DeleteFollow(idUserPerforming int, idUserToUnFollow int) error {

	_, err := db.c.Exec(`
	DELETE 
	FROM Follow
	WHERE IdUserFollowed=? and IdUser=?`, idUserToUnFollow, idUserPerforming)
	if err != nil {
		return err
	}
	return err
}

func (db *appdbimpl) GetFollowersList(idUser int) ([]string, error) {
	var FollowersList []string

	FollowerRows, err := db.c.Query(`
	SELECT Username
	FROM Follow, User
	WHERE Follow.IdUserFollowed=? `, idUser)

	if err != nil {
		return FollowersList, err
	}

	defer FollowerRows.Close()

	for FollowerRows.Next() {
		var name string
		err = FollowerRows.Scan(&name)
		if err != nil {
			return FollowersList, err
		}
		FollowersList = append(FollowersList, name)
	}
	if FollowerRows.Err() != nil {
		return FollowersList, err
	}

	// format the list
	return FollowersList, err

}

func (db *appdbimpl) GetFollowedList(idUser int) ([]string, error) {
	var FollowersList []string

	FollowerRows, err := db.c.Query(`
	SELECT Username
	FROM Follow, User
	WHERE Follow.IdUser=? `, idUser)

	if err != nil {
		return FollowersList, err
	}

	defer FollowerRows.Close()

	for FollowerRows.Next() {
		var name string
		err = FollowerRows.Scan(&name)
		if err != nil {
			return FollowersList, err
		}
		FollowersList = append(FollowersList, name)
	}

	if FollowerRows.Err() != nil {
		return FollowersList, err
	}

	// format the list
	return FollowersList, err
}

func (db *appdbimpl) GetFollowersNumber(idUser int) (int, error) {

	var followersNumber int
	err := db.c.QueryRow(`
	SELECT COUNT(*)
	FROM Follow
	WHERE  Follow.IdUserFollowed=? `, idUser).Scan(&followersNumber)

	if err != nil {
		return 0, err
	}

	return followersNumber, err
}

func (db *appdbimpl) GetFollowedNumber(idUser int) (int, error) {
	var followedNumber int
	err := db.c.QueryRow(`
	SELECT COUNT(*)
	FROM Follow
	WHERE Follow.IdUser=? `, idUser).Scan(&followedNumber)

	if err != nil {
		return 0, err
	}
	return followedNumber, err
}

func (db *appdbimpl) GetFollowState(userPerformingId int, profileOwnerId int) (bool, error) {
	var followState bool
	err := db.c.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM Follow
			WHERE IdUser=?
			AND IdUserFollowed=?)
			`, userPerformingId, profileOwnerId).Scan(&followState)

	if err != nil {
		return false, err
	}

	return followState, err
}
