package pkg

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

func UnpackingString(s string) (string, error) {
	var resultString bytes.Buffer
	stringElements := []rune(s)

	if !isCorrectString(stringElements) {
		return "", errors.New("Некорректная строка")
	}

	for i := 0; i < len(stringElements); i++ {
		stringElement := string(stringElements[i])
		if stringElement == `\` {
			if i < (len(stringElements)-2) && unicode.IsDigit(stringElements[i+2]) {
				nextStringElement := string(stringElements[i+1])
				index := i + 2
				number := getNumber(stringElements, &index)
				for j := 0; j < number; j++ {
					resultString.WriteString(nextStringElement)
				}
				i = index
			} else {
				continue
			}
		} else if i < len(stringElements)-1 && unicode.IsDigit(stringElements[i+1]) {
			index := i + 1
			number := getNumber(stringElements, &index)
			for j := 0; j < number; j++ {
				resultString.WriteString(stringElement)
			}
			i = index - 1
		} else {
			resultString.WriteString(stringElement)
		}
	}
	return resultString.String(), nil
}

func getNumber(sliceRunesString []rune, i *int) int {
	var num int
	for unicode.IsDigit(sliceRunesString[*i]) {
		k, err := strconv.Atoi(string(sliceRunesString[*i]))
		if err != nil {
			break
		}
		num = num*10 + k
		*i++
		if *i >= len(sliceRunesString) {
			break
		}
	}
	return num
}

func isCorrectString(sliceRunesString []rune) bool {
	if len(sliceRunesString) == 0 {
		return true
	}
	for _, v := range sliceRunesString {
		if string(v) == `\` || unicode.IsLetter(v) {

			return true
		}
	}
	return false
}
