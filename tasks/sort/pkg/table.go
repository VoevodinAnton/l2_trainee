package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type table struct {
	textByLines []string
}

func InitTable(lines []string) (*table, error) {
	if lines == nil {
		return nil, errors.New("empty text")
	}
	return &table{lines}, nil
}

func (table *table) writeToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, line := range table.textByLines {
		if table.textByLines[len(table.textByLines)-1] == line {
			fmt.Fprint(f, line)
			break
		}
		fmt.Fprintln(f, line)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadFromFile(in io.Reader) (*table, error) {
	var lines []string
	reader := bufio.NewReader(in)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return &table{lines}, nil
}
