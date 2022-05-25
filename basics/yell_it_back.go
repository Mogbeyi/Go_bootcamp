package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	repeat := strings.Repeat("!", l)

	s:= repeat + msg + repeat
	s = strings.ToUpper(s)

	fmt.Println(s)
}
