package main

import (
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Initializing REST server with Go...")
	routes.HandleRequest()
}
