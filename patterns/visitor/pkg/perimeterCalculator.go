package pkg

import (
	"fmt"
	"math"
)

type PerimeterCalculator struct {
	perimeter int
}

func (p *PerimeterCalculator) visitorForSquare(s *Square) {
	p.perimeter = 4 * s.Side
	fmt.Println("Calculating perimeter for square:", p.perimeter)
}

func (p *PerimeterCalculator) visitorForCircle(c *Circle) {
	p.perimeter = int(float32(2) * math.Pi * float32(c.Radius))
	fmt.Println("Calculating perimeter for circle:", p.perimeter)
}

func (p *PerimeterCalculator) visitorForRectangle(r *Rectangle) {
	p.perimeter = 2*r.Length + 2*r.Width
	fmt.Println("Calculating perimeter for rectangle:", p.perimeter)
}
