// Drop in replacement for the os/user package. Useful when building without cgo.
package userx

import (
	"errors"
	"os/user"
)

var InvalidUserDatabaseError = errors.New("invalid user database")

// Current returns the current user.
func Current() (*user.User, error) {
	return current()
}

// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type UnknownUserError.
func Lookup(username string) (*user.User, error) {
	return lookup(username)
}

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type UnknownUserIdError.
func LookupId(uid string) (*user.User, error) {
	return lookupId(uid)
}
