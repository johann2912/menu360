package main

import (
	"fmt"
	"log"
	"menu360/config"
)

func main() {
	fmt.Println("Hello, World!")

	config.LoadEnv()

	config.ConnectDB()
}
