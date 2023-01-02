package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		panic("Error: Missing file name")
	}

	conn, err := net.Dial("tcp", "192.168.1.72:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	file, err := os.Open(argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file data and write it to the connection
	io.Copy(conn, file)
}
