package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Greper struct {
	textByLines        []string
	flags              []string
	targetString       string
	searchedLines      []string
	countOfLinesBefore int
	countOfLinesAfter  int
	ignoreCase         bool
	lineNumber         bool
	count              int
	needToCount        bool
	invert             bool
	regular            bool
}

func InitGreper(lines []string, fs []string, targetString string) (*Greper, error) {
	if lines == nil {
		return nil, errors.New("empty text")
	}
	return &Greper{
		textByLines:  lines,
		flags:        fs,
		targetString: targetString,
		regular:      true,
	}, nil
}

func (g *Greper) GetResult() string {
	if g.needToCount {
		fmt.Println(g.needToCount)
		return strconv.Itoa(g.count)
	} else {
		return strings.Join(g.searchedLines, "\n")
	}
}

func (g *Greper) Start() error {
	err := g.parseFlags()
	if err != nil {
		return err
	}

	g.searchedLines, g.count, err = grepLinesNearSearched(g.textByLines, g.targetString, g.countOfLinesBefore, g.countOfLinesAfter,
		g.ignoreCase, g.lineNumber, g.regular, g.invert)
	return nil
}

func (g *Greper) parseFlags() error {
	for i := 0; i < len(g.flags); i++ {
		flag := g.flags[i]
		switch flag {
		case "-n":
			g.lineNumber = true

		case "-i":
			g.ignoreCase = true

		case "-c":
			g.needToCount = true

		case "-v":
			g.invert = true

		case "-F":
			g.regular = false

		case "-C":
			if i+1 >= len(g.flags) {
				return errors.New("input count of lines for -C")
			}
			i++
			count, err := strconv.Atoi(g.flags[i])
			if err != nil {
				return err
			}
			if g.countOfLinesBefore == 0 {
				g.countOfLinesBefore = count
			}

			if g.countOfLinesAfter == 0 {
				g.countOfLinesAfter = count
			}

		case "-A":
			if i+1 >= len(g.flags) {
				return errors.New("input count of lines for -A")
			}
			i++
			count, err := strconv.Atoi(g.flags[i])
			if err != nil {
				return err
			}
			g.countOfLinesAfter = count

		case "-B":
			if i+1 >= len(g.flags) {
				return errors.New("input count of lines for -B")
			}
			i++
			count, err := strconv.Atoi(g.flags[i])
			if err != nil {
				return err
			}
			g.countOfLinesBefore = count

		default:
			return errors.New("unknown flag: " + g.flags[i])
		}
	}
	return nil
}
