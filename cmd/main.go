package main

import (
	"fmt"
	"menu360/config"
)

func main() {
	fmt.Println("Hello, World!")
	config.LoadEnv()

	config.ConnectDB()
}
