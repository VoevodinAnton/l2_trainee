package pkg

import (
	"fmt"
	"github.com/opesun/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

func getResponse(url string) (*http.Response, error) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func parseSite(site string, filename string) {
	siteHTML, _ := goquery.ParseUrl(site)
	writeFile(filename, strings.NewReader(siteHTML.Html()))
	parseResources(siteHTML)
}

func parseResources(site goquery.Nodes) {
	for _, url := range site.Find("").Attrs("src") {
		fmt.Println(url)
		urlParts := strings.Split(url, "/")
		filename := urlParts[len(urlParts)-1]
		downloadResources(filename, url)
	}
}

func downloadResources(filepath string, url string) error {

	resourse, err := http.Get(url)
	if err != nil {
		return err
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resourse.Body)
	return err
}
