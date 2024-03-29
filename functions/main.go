// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func main() {
// 	p := newParser()

// 	in := bufio.NewScanner(os.Stdin)

// 	for in.Scan() {
// 		p.lines++

// 		parsed, err := parse(p, in.Text())
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		domain, visits := parsed.domain, parsed.visits

// 		// Collect the unique domains
// 		if _, ok := p.sum[domain]; !ok {
// 			p.domains = append(p.domains, domain)
// 		}

// 		// Keep track of total and per domain visits
// 		p.total += visits

// 		// create and assign a new copy of `visit`
// 		p.sum[domain] = result{
// 			domain: domain,
// 			visits: visits + p.sum[domain].visits,
// 		}
// 	}

// 	// Print the visits per domain
// 	sort.Strings(p.domains)

// 	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
// 	fmt.Println(strings.Repeat("-", 45))

// 	for _, domain := range p.domains {
// 		parsed := p.sum[domain]
// 		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
// 	}

// 	// Print the total visits for all domains
// 	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

// 	// Let's handle the error
// 	if err := in.Err(); err != nil {
// 		fmt.Println("> Err:", err)
// 	}
// }
package main

import "fmt"

func main() {
	stats := map[int]int{1: 10, 10: 2}
	incrAll(stats)
	fmt.Println(stats)
}

func incrAll(stats map[int]int) {
	for k := range stats {
		stats[k]++
	}
}