package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	logs := make(map[string]int)

	for in.Scan() {
		result := strings.Split(in.Text(), " ")
		site := result[0]
		count, _ := strconv.Atoi(result[1])
		logs[site] += count
	}

	fmt.Println(logs)
}