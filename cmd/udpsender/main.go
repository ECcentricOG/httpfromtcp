package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("error ", "error ", err)
	}
	conn, err := net.DialUDP(addr.Network(), nil, addr) 
	if err != nil {
		log.Fatal("error ", "error ", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error  ", "error  ", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		if len(input) == 0 {
			continue
		}
		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Fatal("error ", "error ", err)
		}
	}
}
