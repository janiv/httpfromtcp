package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Yo something went wrong opening the file")
	}
	bts := make([]byte, 8)

	curr := ""
	for {
		count, err := file.Read(bts)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		stuff := string(bts[:count])
		curr += stuff
		if strings.Index(curr, "\n") > 0 {
			cutoff := strings.Index(curr, "\n")
			printable := curr[:cutoff]
			if len(strings.TrimSpace(printable)) > 0 {
				fmt.Printf("read: %s\n", printable)
			}
			curr = curr[cutoff+1:]
		}
	}
	if len(strings.TrimSpace(curr)) > 0 {
		fmt.Printf("read: %s\n", curr)
	}

}
