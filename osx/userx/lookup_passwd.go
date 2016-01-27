// +build !windows,!darwin

package userx

import (
	"bufio"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

const UserDatabase = "/etc/passwd"

type PasswdRecord struct {
	Name      string
	Password  string
	Uid       string
	Gid       string
	Gecos     string
	Directory string
	Shell     string
}

func current() (*user.User, error) {
	return lookupPasswd(syscall.Getuid(), "", false)
}

func lookup(username string) (*user.User, error) {
	return lookupPasswd(-1, username, true)
}

func lookupId(uid string) (*user.User, error) {
	i, err := strconv.Atoi(uid)
	if err != nil {
		return nil, err
	}
	return lookupPasswd(i, "", false)
}

func lookupPasswd(uid int, username string, lookupByName bool) (*user.User, error) {
	f, err := os.Open(UserDatabase)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	exist := false
	p := &PasswdRecord{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.SplitN(scanner.Text(), ":", 7)
		if strings.HasPrefix(fields[0], "#") {
			continue
		}
		if len(fields) < 7 {
			return nil, InvalidUserDatabaseError
		}
		p = &PasswdRecord{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5], fields[6]}
		if lookupByName {
			if username == p.Name {
				exist = true
				break
			}
		} else {
			if strconv.Itoa(uid) == p.Uid {
				exist = true
				break
			}
		}
	}
	if !exist {
		if lookupByName {
			return nil, user.UnknownUserError(username)
		} else {
			return nil, user.UnknownUserIdError(uid)
		}
	}

	u := &user.User{
		Username: p.Name,
		Uid:      p.Uid,
		Gid:      p.Gid,
		Name:     p.Gecos,
		HomeDir:  p.Directory,
	}
	// The gecos field isn't quite standardized.  Some docs
	// say: "It is expected to be a comma separated list of
	// personal data where the first item is the full name of the
	// user."
	if i := strings.Index(u.Name, ","); i >= 0 {
		u.Name = u.Name[:i]
	}
	return u, nil
}
