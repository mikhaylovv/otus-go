package main

import (
	"log"
	"os"

	"github.com/mikhaylovv/otus-go/hw_7/hw7"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("Minimum arguments count is 3\n" +
			"Usage: go-envdir /path/to/evndir command arg1 arg2")
	}

	env, err := hw7.ReadDir(args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	err = hw7.RunCmd(args[2:], env)
	if err != nil {
		log.Fatal(err.Error())
	}
}
