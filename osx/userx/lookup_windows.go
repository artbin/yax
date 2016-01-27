package userx

import (
	"os"
	"os/user"
)

func current() (*user.User, error) {
	return user.Current()
}

func lookup(username string) (*user.User, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return nil, err
	}
	u.HomeDir = detectHomeDir(username)
	return u, nil
}

func lookupId(uid string) (*user.User, error) {
	u, err := user.LookupId(uid)
	if err != nil {
		return nil, err
	}
	u.HomeDir = detectHomeDir(u.Name)
	return u, nil
}

func detectHomeDir(username string) string {
	dir := `C:\Users\` + username
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dir = `C:\Documents and Settings\` + username
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return ""
		}
	}
	return dir
}
