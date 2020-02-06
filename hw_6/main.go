package main

import (
	"flag"
	"log"

	"github.com/mikhaylovv/otus-go/hw_6/hw6"
)

var (
	from string
	to string
	limit int64
	offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write copy of \"From\" file")
	flag.Int64Var(&limit, "limit", 0, "max number of bytes to copy, by default copy all file")
	flag.Int64Var(&offset, "offset", 0, "offset in \"From\" file")
}

func main() {
	flag.Parse()

	err := hw6.Copy(from, to, limit, offset)

	if err != nil {
		log.Fatal(err.Error())
	}
}
