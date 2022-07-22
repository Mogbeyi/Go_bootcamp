package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords) // purpose of this line is to split the text into individual words
	total, words := 0, make(map[string]int)

	for in.Scan() {
		total++
		words[in.Text()]++
	}
	fmt.Println(words)

	fmt.Printf("There are %d words, %d of them are unique.\n",total, len(words))
}