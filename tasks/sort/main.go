package main

import (
	"flag"
	"os"
	"sort/pkg"
)

func main() {
	file, err := os.Open(flag.Lookup("f").Value.String())
	if err != nil {
		println(err.Error())
		return
	}
	table, err := pkg.ReadFromFile(file)
	if err != nil {
		println(err.Error())
		return
	}

	table.Sort()
}
