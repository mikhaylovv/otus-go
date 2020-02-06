package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestCopyProg(t *testing.T) {
	cmd := exec.Command("sh", "test.sh")
	cmd.Dir = "test"
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}