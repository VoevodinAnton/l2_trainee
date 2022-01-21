package pkg

import (
	"flag"
	"fmt"
)

var (
	fileName                   = flag.String("f", "", "filename")
	targetColumn               = flag.Int("k", 1, "target column")
	sortByNumber               = flag.Bool("n", false, "sort by number")
	shouldReverse              = flag.Bool("r", false, "sort in reverse order")
	shouldUnique               = flag.Bool("u", false, "get unique elements")
	sortByMonth                = flag.Bool("M", false, "sort by month")
	shouldIgnoreTrailingSpaces = flag.Bool("b", false, "ignore trailing spaces ")
	checkSorted                = flag.Bool("c", false, "check sorted data")
	sortByNumberWithSuffix     = flag.Bool("h", false, "sort by number with suffix")
)

func init() {
	flag.Parse()
}

func (table *table) Sort() {

	if *checkSorted {
		if equalSlices(table.textByLines, simpleSort(table.textByLines, *targetColumn, *shouldReverse)) {
			fmt.Println(*targetColumn, "column is sorted")
		} else {
			fmt.Println(*targetColumn, "column is not sorted")
		}
	}

	if *shouldUnique {
		linesWithoutDuplicates := deleteOfDuplicates(table.textByLines)
		table.textByLines = linesWithoutDuplicates
	}
	if *shouldIgnoreTrailingSpaces {

	}

	if *sortByNumber {
		sortedLines := sortByNumberN(table.textByLines, *targetColumn, *shouldReverse)
		table.textByLines = sortedLines
	} else if *sortByMonth {
		sortedLines := sortByMonthM(table.textByLines, *targetColumn, *shouldReverse)
		table.textByLines = sortedLines
	} else if *sortByNumberWithSuffix {
		sortedLines := sortByNumberWithSuffixH(table.textByLines, *targetColumn, *shouldReverse)
		table.textByLines = sortedLines
	}

	table.writeToFile(*fileName)
}
