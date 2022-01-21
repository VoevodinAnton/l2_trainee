package main

import (
	"interaction-with-os/pkg"
	"os"
)

func main() {
	commander, err := pkg.InitSheller(os.Stdout, os.Stdin)
	if err != nil {
		println(err.Error())
	}

	if err := commander.Start(); err != nil {
		println(err.Error())
	}
}
