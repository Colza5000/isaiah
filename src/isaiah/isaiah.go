package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	docker, lookErr := exec.LookPath("docker-compose")
	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println("Go to http://localhost:8080")

	args := []string{"docker-compose", "-f", "../docker/docker-compose.yml", "up"}

	env := os.Environ()

	execErr := syscall.Exec(docker, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
