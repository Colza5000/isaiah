package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	docker, lookErr := exec.LookPath("docker-compose")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"docker-compose", "up"}

	env := os.Environ()

	execErr := syscall.Exec(docker, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
