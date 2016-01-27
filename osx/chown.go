package osx

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Recursive chown
func Chown(name string, uid, gid int) (err error) {
	err = os.Chown(name, uid, gid)
	if err != nil {
		return
	}
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	if fi.IsDir() {
		files, err := ioutil.ReadDir(name)
		if err != nil {
			return err
		}
		for _, fi := range files {
			p := filepath.Join(name, fi.Name())
			err = Chown(p, uid, gid)
			if err != nil {
				return err
			}
		}
	}
	return
}
