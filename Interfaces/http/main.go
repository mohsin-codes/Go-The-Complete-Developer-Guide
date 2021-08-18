package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

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

	//Implenting Custom Writer Interface that logs out elements from the body
	lw := logWriter{}

	//using write method to print the elements of the body
	io.Copy(lw, resp.Body)
}

//Definition of Write function as custom implementation of Writer Interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Length of processed bytes is : ", len(bs))
	return len(bs), nil
}
