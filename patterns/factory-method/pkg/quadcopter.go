package pkg

type quadcopter struct {
	transport
}

func newQuadcopter() iTransport {
	return &quadcopter{
		transport: transport{
			name:  "Quadcopter",
			speed: 14,
		},
	}
}
