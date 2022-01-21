package pkg

import "fmt"

type FilterFoFirstAlgorithm struct{}

func (f *FilterFoFirstAlgorithm) doSearch(filters map[string]int) {
	fmt.Println("First implements strategy", filters)
}
