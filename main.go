package main

import (
	"fmt"
	"urlshortner/internal/api"
)

func main() {
	api.InitializeApp()
	fmt.Println("Url shortener service is up")
}
