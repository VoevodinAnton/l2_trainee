package pkg

import "strings"

func getElementByColumn(line string, numberColumn int) (string, bool) {
	if line == "" {
		return "", false
	}

	if numberColumn < 0 {
		return "", false
	}

	elements := strings.Fields(line)

	if numberColumn < len(elements) {
		return elements[numberColumn], true
	} else {
		return "", false
	}
}

func deleteOfDuplicates(lines []string) []string {
	var result []string
	keyToUniqMap := make(map[string]bool)

	for _, line := range lines {
		if _, contain := keyToUniqMap[line]; !contain {
			result = append(result, line)
			keyToUniqMap[line] = true
		}
	}

	return result
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
