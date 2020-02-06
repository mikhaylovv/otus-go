package main

import (
	"log"
	"os/exec"
	"testing"
)

func TestCopyProg(t *testing.T) {
	cmd := exec.Command("sh", "test.sh")
	cmd.Dir = "test"

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}