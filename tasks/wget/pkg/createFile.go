package pkg

import (
	"io"
	"net/url"
	"os"
	"strings"
)

func createFile(fullUrl, fileName string) (string, error) {
	if fileName == "" {
		fileUrl, err := url.Parse(fullUrl)
		if err != nil {
			return "", err
		}
		path := fileUrl.Path
		urlParts := strings.Split(path, "/")
		fileName = urlParts[len(urlParts)-1]
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return fileName, nil
}

func writeFile(filename string, resp io.Reader) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()

	io.Copy(file, resp)

	defer file.Close()

	return err
}
