package utils

import (
	"bytes"
	"os"
	"os/exec"
)

func CmdWithDir(dir string, commandName string, params ...string) (string, error) {
	cmd := exec.Command(commandName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return out.String(), err
}

func CmdWithDirNoOut(dir string, commandName string, params ...string) error {
	cmd := exec.Command(commandName, params...)
	cmd.Dir = dir
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}
