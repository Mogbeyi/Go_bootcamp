package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Give me something to mask!")
		return 
	}

	const (
		link = "https://"
		nlink = len(link)
	)

	var (
		text = args[0]
		size = len(text)
		buf = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		buf = append(buf, text[i])

		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			fmt.Println(text[i : i + nlink])
		}
	}
}