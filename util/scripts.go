package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func ExecuteCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := io.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
