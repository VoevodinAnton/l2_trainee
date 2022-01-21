package pkg

import (
	"fmt"
	"math"
)

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitorForSquare(s *Square) {
	a.area = s.Side * s.Side
	fmt.Println("Calculating area for square:", a.area)
}

func (a *AreaCalculator) visitorForCircle(s *Circle) {
	a.area = int(math.Pi * float64(s.Radius) * float64(s.Radius))
	fmt.Println("Calculating area for circle:", a.area)
}

func (a *AreaCalculator) visitorForRectangle(s *Rectangle) {
	a.area = s.Width * s.Length
	fmt.Println("Calculating area for rectangle:", a.area)
}
