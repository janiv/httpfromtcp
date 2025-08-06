package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:42069")
	if err != nil {
		fmt.Println("net listen failed")
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("where connection")
			break
		}
		fmt.Println("connected")
		go func(c net.Conn) {
			linesChan := getLinesChannel(c)
			for line := range linesChan {
				fmt.Println(line)
			}
			fmt.Println("Message received closing channel")
		}(conn)
	}
	fmt.Println("Program terminated, closing channel")
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	strChannel := make(chan string)
	go func() {
		defer f.Close()
		defer close(strChannel)
		curr := ""
		for {
			bts := make([]byte, 8)
			count, err := f.Read(bts)
			if err != nil {
				if curr != "" {
					strChannel <- curr
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			stuff := string(bts[:count])
			parts := strings.Split(stuff, "\n")
			for i := 0; i < len(parts)-1; i++ {
				strChannel <- fmt.Sprintf("%s%s", curr, parts[i])
				curr = ""
			}
			curr += parts[len(parts)-1]
		}
		if len(strings.TrimSpace(curr)) > 0 {
			strChannel <- curr
		}
	}()
	return strChannel
}
