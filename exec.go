package main

import (
	"os/exec"
)

func IsExecutable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	err := cmd.Run()
	return err == nil // lazy but fine for now
}
