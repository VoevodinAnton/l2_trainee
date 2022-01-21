package main

import (
	"fmt"
	"unpacking/pkg"
)

func main() {
	slice := []string{"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`, `a0b2c12`, `q\413`, `2352352332`, `\24`}
	for _, s := range slice {
		res, err := pkg.UnpackingString(s)
		if err != nil {
			fmt.Printf("%v => \" \" (%v)\n", s, err)
		} else {
			fmt.Printf("%v => %v\n", s, res)
		}
	}
}
