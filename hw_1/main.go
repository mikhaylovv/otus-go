package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	if time, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(time)
	}
}
