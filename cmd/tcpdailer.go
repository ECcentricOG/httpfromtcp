package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:42069")
	if err != nil {
		log.Fatal("Error", err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")	
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error", err)
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Fatal("error", err)
		}
	}
}
