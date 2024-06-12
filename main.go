package main

import (
	"fmt"
	"github.com/shariq/golang-dev-logic-challenge-shariqcheema/routes"
)

func main() {
	router := routes.SetupRouter()

	fmt.Println("Starting server on port 8080")
	router.Run(":8080")
}
