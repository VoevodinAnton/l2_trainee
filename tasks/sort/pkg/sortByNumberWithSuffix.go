package pkg

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

var suffixes = map[string]int{
	"KB": 1 << 10,
	"K":  1 << 10,
	"MB": 1 << 20,
	"M":  1 << 20,
	"GB": 1 << 30,
	"G":  1 << 30,
	"TB": 1 << 40,
	"T":  1 << 40,
}

func parseStringWithSuffixInNumber(s string) (float64, error) {
	s = strings.ToUpper(s)
	for key, val := range suffixes {
		i := strings.Index(s, key)
		if i > -1 {
			numberString := s[:i]
			number, err := strconv.ParseFloat(numberString, 64)
			if err != nil {
				return 0, err
			}
			numberInBytes := number * float64(val)
			return numberInBytes, nil
		}
	}
	return 0, errors.New("wrong number")
}

func sortByNumberWithSuffixH(lines []string, numberColumn int, shouldReverse bool) []string {
	notNumberLines := make([]string, 0, len(lines))
	targetLines := make(map[float64][]string)
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		element, found := getElementByColumn(line, numberColumn-1)
		if !found {
			notNumberLines = append(notNumberLines, line)
		} else {
			key, err := parseStringWithSuffixInNumber(element)
			if err != nil {
				notNumberLines = append(notNumberLines, line)
			} else {
				targetLines[key] = append(targetLines[key], line)
			}
		}
	}

	keys := make([]float64, 0, len(targetLines))
	for k := range targetLines {
		keys = append(keys, k)
	}

	if shouldReverse {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		sort.Sort(sort.Reverse(sort.StringSlice(notNumberLines)))

		for _, key := range keys {
			result = append(result, targetLines[key]...)
		}
		result = append(result, notNumberLines...)
	} else {
		sort.Strings(notNumberLines)
		sort.Float64s(keys)

		for _, key := range keys {
			result = append(result, targetLines[key]...)
		}
		result = append(result, notNumberLines...)
	}

	return result
}
