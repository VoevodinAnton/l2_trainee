package pkg

import (
	"sort"
	"strconv"
)

func sortByNumberN(lines []string, numberColumn int, shouldReverse bool) []string {
	notNumberLines := make([]string, 0, len(lines))
	targetLines := make(map[int][]string)
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		element, _ := getElementByColumn(line, numberColumn-1)
		if _, err := strconv.Atoi(element); err != nil {
			notNumberLines = append(notNumberLines, line)
		} else {
			key, _ := strconv.Atoi(element)
			targetLines[key] = append(targetLines[key], line)
		}
	}

	keys := make([]int, 0, len(targetLines))
	for k := range targetLines {
		keys = append(keys, k)
	}

	if shouldReverse {
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	} else {
		sort.Ints(keys)
	}

	for _, key := range keys {
		result = append(result, targetLines[key]...)
	}
	result = append(result, notNumberLines...)
	return result
}
