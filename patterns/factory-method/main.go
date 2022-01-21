package main

import (
	"factory/pkg"
	"fmt"
)

func main() {
	scooter, _ := pkg.GetTransport("scooter")
	quad, _ := pkg.GetTransport("quadcopter")

	fmt.Println(scooter)
	fmt.Println(quad)
}
