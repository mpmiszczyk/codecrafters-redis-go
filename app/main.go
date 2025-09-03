package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using default values")
	}

	// Get port from environment variable or use default
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Failed to bind to port %s\n", port)
		os.Exit(1)
	}
	fmt.Println("before Accept()")
	conn, err := l.Accept()
	fmt.Println("after Accept()")
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("before PONG")
	conn.Write([]byte("+PONG\r\n"))
	fmt.Println("after PONG")
}
