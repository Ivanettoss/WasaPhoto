package database

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"

func (db *appdbimpl) GetStream(userPerformingId int, userPerformingName string) (components.Stream, error) {
	var stream components.Stream
	var photos []components.Photo

	StreamRows, err := db.c.Query(`
	SELECT User.Username,Photo.IdUser,Photo.IdPhoto,Photo.DateTime,Photo.Path
	FROM User,Photo
	WHERE  
	User.IdUser=Photo.IdUser 
	and User.IdUser=Follow.IdUserFollowed and Follow.IdUser=? 
	and ? NOT IN (SELECT IdUserBanned FROM Ban WHERE Ban.IdUser=User.IdUser)
	and User.IdUser NOT IN ( SELECT IdUserBanned FROM Ban WHERE Ban.IdUser=?)
	ORDER BY Photo.DateTime
	`, userPerformingId, userPerformingId, userPerformingId)

	if err != nil {
		return stream, err
	}

	for StreamRows.Next() {
		var photo components.Photo
		StreamRows.Scan(&photo.Username, &photo.IdPhoto, &photo.UploadDataTime, &photo.PhotoBytes)
		photos = append(photos, photo)
	}

	stream.Photos = photos
	stream.Username = userPerformingName

	return stream, nil
}
