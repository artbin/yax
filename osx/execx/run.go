package execx

import (
	"log"
	"os"
	"os/exec"
)

func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func RunOrExit(name string, args ...string) {
	if err := Run(name, args...); err != nil {
		log.Fatalln(err)
	}
}
