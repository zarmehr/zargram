package models

import "fmt"

var (
	// ErrUserNotFound ...
	ErrUserNotFound = fmt.Errorf("user not found")

	// ErrInternalServer ...
	ErrInternalServer = fmt.Errorf("internal server error")

	// ErrUserAlreadyHasAvatar ...
	ErrUserAlreadyHasAvatar = fmt.Errorf("user already has an avatar")

	// ErrFailToRetrieve...
	ErrFailToRetrieve = fmt.Errorf("failed to retrieve file")

	//ErrInvalidId...
	ErrInvalidId = fmt.Errorf("failed to retrieve id")
	//
	////ErrBadRequest...
	//ErrBadRequest =fmt.Errorf("Bad request")
)
