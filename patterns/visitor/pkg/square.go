package pkg

type Square struct {
	Side int
}

func (s *Square) Accept(v visitor) {
	v.visitorForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}
