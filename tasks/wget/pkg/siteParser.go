package pkg

import (
	"fmt"
	"github.com/opesun/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

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
