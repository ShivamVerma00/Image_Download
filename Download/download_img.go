package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func Download_Images(images []string) string {

	wg.Add(len(images))

	limit := make(chan int, 10) // channel for 10
	defer close(limit)
	for _, images := range images {
		limit <- 1
		go func(images string) {
			defer wg.Done()

			tokens := strings.Split(images, "/")

			imageName := tokens[len(tokens)-1]

			u, err := url.Parse(images)
			if err != nil {
				panic(err)
			}

			if u.Scheme == "https" {
				output, err := os.Create(imageName)
				if err != nil {
					log.Fatal(err)
				}
				defer output.Close()

				res, err := http.Get(images)
				if err != nil {
					log.Fatal(err)
				} else {
					defer res.Body.Close()
					_, err = io.Copy(output, res.Body)
					if err != nil {
						log.Fatal(err)
					} else {
						fmt.Println("Downloaded", imageName)
					}
				}
			}
			<-limit

		}(images)
	}
	wg.Wait()
	return "done"
}
