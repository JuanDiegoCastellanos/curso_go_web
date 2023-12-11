package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	rta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	e := escritorWeb{}

	io.Copy(e, rta.Body)

}
