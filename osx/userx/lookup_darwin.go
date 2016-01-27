package userx

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	"github.com/DHowett/go-plist"
)

type userplist struct {
	HomeDir []string `plist:"dsAttrTypeStandard:NFSHomeDirectory"`
	Gid     []string `plist:"dsAttrTypeStandard:PrimaryGroupID"`
	Uid     []string `plist:"dsAttrTypeStandard:UniqueID"`
	Name    []string `plist:"dsAttrTypeStandard:RealName"`
}

func current() (*user.User, error) {
	return lookupDscl(syscall.Getuid(), "", false)
}

func lookup(username string) (*user.User, error) {
	return lookupDscl(-1, username, true)
}

func lookupId(uid string) (*user.User, error) {
	i, err := strconv.Atoi(uid)
	if err != nil {
		return nil, err
	}
	return lookupDscl(i, "", false)
}

func lookupDscl(uid int, username string, lookupByName bool) (*user.User, error) {
	dscl := exec.Command("/usr/bin/dscl", ".", "-list", "users", "uid")

	dsclbuf, err := dscl.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = dscl.Start(); err != nil {
		return nil, err
	}

	name := ""
	id := ""

	exist := false
	scanner := bufio.NewScanner(dsclbuf)
	for scanner.Scan() {
		scannerWord := bufio.NewScanner(bytes.NewReader(scanner.Bytes()))
		scannerWord.Split(bufio.ScanWords)
		scannerWord.Scan()
		name = scannerWord.Text()
		scannerWord.Scan()
		id = scannerWord.Text()

		if lookupByName {
			if username != name && "_"+username != name {
				continue
			}
		} else {
			if strconv.Itoa(uid) != id {
				continue
			}
		}

		dscl = exec.Command("/usr/bin/dscl", "-plist", ".", "-read", "users/"+name)
		dsclbuf, err = dscl.StdoutPipe()
		if err != nil {
			return nil, err
		}
		if err = dscl.Start(); err != nil {
			return nil, err
		}

		var uplist userplist
		buf, err := ioutil.ReadAll(dsclbuf)
		if err != nil {
			return nil, err
		}
		_, err = plist.Unmarshal(buf, &uplist)
		if err != nil {
			return nil, err
		}
		exist = true

		if name[0] == '_' {
			name = name[1:]
		}

		u := &user.User{
			Username: name,
			Uid:      id,
			Gid:      uplist.Gid[0],
			Name:     uplist.Name[0],
			HomeDir:  uplist.HomeDir[0],
		}
		return u, nil
	}
	if !exist {
		if lookupByName {
			return nil, user.UnknownUserError(username)
		} else {
			return nil, user.UnknownUserIdError(uid)
		}
	}
	return nil, InvalidUserDatabaseError
}
