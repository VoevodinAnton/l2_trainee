package main

import (
	"anagram/pkg"
	"fmt"
)

func main() {
	slice := []string{
		"пятак",
		"листок",
		"олег",
		"пятка",
		"тяпка",
		"кот",
		"ток",
		"пот",
		"топ",
		"ЛЕГО",
		"слиток",
		"мукА",
		"столик",
		"акума",
	}

	result := pkg.SearchAnagramByDictionary(&slice)
	fmt.Println(*result)
}
