package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	address, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Println("udp resolve address failed")
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing UDP: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Sending to %s. Type your yap and press enter to send.\n", "localhost:42069")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v]\n", err)
			os.Exit(1)
		}
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error sending message: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Message sent: %s", message)
	}

}
