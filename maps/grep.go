package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, args[0]) {
			fmt.Println(in.Text())
		}
	}
}