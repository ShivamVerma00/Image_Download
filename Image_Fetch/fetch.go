package Image_Fetch

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"image_download/download"

	"golang.org/x/net/html"
)

func Fetch_Image(url string) string {
	result := make([]string, 0)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	html1, err := ioutil.ReadAll(resp.Body) // reads html as slice
	if err != nil {
		log.Fatal(err)
	}
	html2 := string(html1) // convert slice of bytes to string

	doc, err := html.Parse(strings.NewReader(html2)) //parsing
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" { //fetching images
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = append(result, img.Val)

				}
			}

		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(doc)

	images := result

	msg := download.Download_Images(images)
	return msg

}
