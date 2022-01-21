package pkg

type electricScooter struct {
	transport
}

func newElectricScooter() iTransport {
	return &electricScooter{
		transport: transport{
			name:  "Scooter",
			speed: 4,
		},
	}
}
