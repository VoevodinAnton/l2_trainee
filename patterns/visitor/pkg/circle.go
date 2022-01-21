package pkg

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v visitor) {
	v.visitorForCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}
