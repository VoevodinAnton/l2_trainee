package pkg

type Rectangle struct {
	Length int
	Width  int
}

func (t *Rectangle) Accept(v visitor) {
	v.visitorForRectangle(t)
}

func (t *Rectangle) GetType() string {
	return "Rectangle"
}
