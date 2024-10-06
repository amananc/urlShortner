package main

import (
	"fmt"
	"urlshortner/config"
)

func main() {
	fmt.Println("Url shortner service is up")
	config.InitializeApp()
}