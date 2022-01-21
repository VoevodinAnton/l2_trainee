package pkg

import (
	"sort"
	"strings"
)

type Month []string

func (m Month) Len() int           { return len(m) }
func (m Month) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Month) Less(i, j int) bool { return month[m[i]] < month[m[j]] }

var month = map[string]int{
	"JAN": 1,
	"FAB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

func sortByMonthM(lines []string, numberColumn int, shouldReverse bool) []string {
	notMonthLines := make([]string, 0, len(lines))
	targetLines := make(map[string][]string)
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		element, found := getElementByColumn(line, numberColumn-1)
		if !found {
			notMonthLines = append(notMonthLines, line)
		} else {
			key := strings.ToUpper(element)
			if _, ok := month[key]; !ok {
				notMonthLines = append(notMonthLines, line)
			} else {
				targetLines[key] = append(targetLines[key], line)
			}
		}
	}

	keys := make(Month, 0, len(targetLines))
	for k := range targetLines {
		keys = append(keys, k)
	}

	if shouldReverse {
		sort.Sort(sort.Reverse(keys))
		sort.Sort(sort.Reverse(sort.StringSlice(notMonthLines)))

		for _, key := range keys {
			result = append(result, targetLines[key]...)
		}
		result = append(result, notMonthLines...)
	} else {
		sort.Strings(notMonthLines)
		sort.Sort(keys)

		for _, key := range keys {
			result = append(result, targetLines[key]...)
		}
		result = append(result, notMonthLines...)
	}

	return result
}
