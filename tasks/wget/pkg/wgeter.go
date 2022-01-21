package pkg

import (
	"errors"
)

type Wgeter struct {
	url      string
	fileName string
}

func InitWgeter(url string, filename string) (*Wgeter, error) {
	if url == "" {
		return nil, errors.New("empty text")
	}
	return &Wgeter{url, filename}, nil
}

func (w *Wgeter) Start() error {
	fileName, err := createFile(w.url, w.fileName)
	if err != nil {
		return err
	}
	w.fileName = fileName

	parseSite(w.url, w.fileName)

	return nil
}
