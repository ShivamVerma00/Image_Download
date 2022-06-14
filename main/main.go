package main

import (
	"fmt"
	Image_Fetch "image_download/image_fetch"
)

func main() {
	/*	fmt.Println("Enter url: ")
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("Invalid...")
		}
		fmt.Println(str)*/

	//	url := string(str)
	url := "https://www.tftus.com"

	fmt.Println(Image_Fetch.Fetch_Image(url))
}
