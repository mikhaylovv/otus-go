package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	if time, error := ntp.Time("0.beevik-ntp.pool.ntp.org"); error == nil {
		fmt.Println(time)
	} else {
		os.Stderr.WriteString(error.Error())
		os.Exit(1)
	}
}
