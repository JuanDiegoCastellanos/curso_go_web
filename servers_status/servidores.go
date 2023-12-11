package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	c := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	for _, servidor := range servidores {
		go revisarServer(servidor, c)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-c)
	}

	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}
	fin := time.Since(inicio)

	fmt.Println("Tiempo tardado: ", fin)

}
func revisarServer(address string, canal chan string) {
	_, err := http.Get(address)
	if err != nil {
		fmt.Println("Servidor no disponible :(")
		canal <- address + "No disponible"
	} else {
		fmt.Println("Server online")
		canal <- address + "Disponible"
	}
}
