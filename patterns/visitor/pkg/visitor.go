package pkg

type visitor interface {
	visitorForSquare(*Square)
	visitorForCircle(*Circle)
	visitorForRectangle(*Rectangle)
}
