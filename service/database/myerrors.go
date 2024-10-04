package database

import (
	"errors"
)

// user
var ErrUserNotFound = errors.New("The user you are looking for was not found")

// photo
var ErrPhotoNotFound = errors.New("The picture you are looking for was not found")

var ErrPageNotFound = errors.New("The page you are looking for was not found")

var ErrLikeNotFound = errors.New("The like you are looking for was not found")

var ErrCommNotFound = errors.New("The comment you are looking for was not found")

var ErrUserBanned = errors.New("User banned")


