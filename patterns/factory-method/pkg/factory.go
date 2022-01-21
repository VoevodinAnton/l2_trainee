package pkg

import "fmt"

func GetTransport(tt string) (iTransport, error) {
	if tt == "scooter" {
		return newElectricScooter(), nil
	}
	if tt == "quadcopter" {
		return newQuadcopter(), nil
	}
	return nil, fmt.Errorf("Wrong type")
}
