package osx

import (
	"io"
	"os"
)

// CopyFile makes a copy of a file preserving its access flags
func CopyFile(dst, src string) (err error) {
	srcf, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcf.Close()
	fi, err := srcf.Stat()
	if err != nil {
		return
	}
	dstf, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, fi.Mode())
	if err != nil {
		return
	}
	defer dstf.Close()
	if _, err = io.Copy(dstf, srcf); err != nil {
		return
	}
	err = dstf.Sync()
	return
}
