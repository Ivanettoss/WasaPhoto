package database

import (
	"errors"
)

var erroreEsempio = errors.New("bro ur not drippin")

// user
var ErrUserNotFound = errors.New("The user you are looking for was not found")

// photo
var ErrPhotoNotFound = errors.New("The picture you are looking for was not found")

var ErrPageNotFound = errors.New("The page you are looking for was not found")

var ErrLikeNotFound = errors.New("The like you are looking for was not found")
