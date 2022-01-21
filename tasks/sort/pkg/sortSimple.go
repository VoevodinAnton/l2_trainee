package pkg

import (
	"sort"
	"strings"
)

func simpleSort(lines []string, numberColumn int, shouldReverse bool) []string {
	targetLines := make(map[string][]string)

	for _, line := range lines {
		element, _ := getElementByColumn(line, numberColumn-1)
		lowerLine := strings.ToLower(element)
		targetLines[lowerLine] = append(targetLines[lowerLine], line)
	}

	keys := make([]string, 0, len(targetLines))
	for k := range targetLines {
		keys = append(keys, k)
	}

	if shouldReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	var result []string
	for _, key := range keys {
		result = append(result, targetLines[key]...)
	}
	return result
}
