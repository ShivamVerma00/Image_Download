package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter url: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Invalid...")
	}

	//	url := string(str)
	//	url := "https://golang.org"
	//	msg := Image_Fetch.Fetch_Image(url)
	//	fmt.Println(msg)
}
