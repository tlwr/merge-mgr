package main

import (
	"os/exec"
)

func IsExecutable(name string) bool {
	cmd := exec.Command("which", name)
	err := cmd.Run()
	return err == nil // lazy but fine for now
}
