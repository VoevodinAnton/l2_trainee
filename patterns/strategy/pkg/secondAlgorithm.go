package pkg

import "fmt"

type FilterFoSecondAlgorithm struct{}

func (f FilterFoSecondAlgorithm) doSearch(filters map[string]int) {
	fmt.Println("First implements strategy", filters)
}
