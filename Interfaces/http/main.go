package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	//using read method to print the elements of the body
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//using write method to print the elements of the body
	io.Copy(os.Stdout, resp.Body)
}
