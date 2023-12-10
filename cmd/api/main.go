package main

import (
	"fmt"
	"htmx-tutorial/internal/server"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.New()

	server.RegisterFiberRoutes()
	port := os.Getenv("PORT")
	fmt.Println(port)
	err := server.Listen(port)
	if err != nil {
		panic("cannot start server")
	}
}
