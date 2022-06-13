package main

import (
	Image_Fetch "image_download/image_fetch"
	"testing"
)

func TestFindImages1(t *testing.T) {

	tst := "https://www.geeksforgeeks.org/"

	expectedOutput := "done"
	output := Image_Fetch.Fetch_Image(tst)

	if output != expectedOutput {
		t.Errorf("got %q, wanted %q", output, expectedOutput)
	}

}
