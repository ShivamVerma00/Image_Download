package main

import (
	"bufio"
	"fmt"
	Image_Fetch "image_download/image_fetch"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter url: ")
	reader := bufio.NewReader(os.Stdin)
	url, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Invalid...")
	}

	msg := Image_Fetch.Fetch_Image(url)
	fmt.Println(msg)
}
