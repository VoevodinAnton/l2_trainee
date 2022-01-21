package main

import "visitor/pkg"

func main() {
	circle := &pkg.Circle{Radius: 4}
	square := &pkg.Square{Side: 1}
	rectangle := &pkg.Rectangle{Length: 8, Width: 10}

	areaCalculator := &pkg.AreaCalculator{}

	circle.Accept(areaCalculator)
	square.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	perimeterCalculator := &pkg.PerimeterCalculator{}

	circle.Accept(perimeterCalculator)
	square.Accept(perimeterCalculator)
	rectangle.Accept(perimeterCalculator)
}
