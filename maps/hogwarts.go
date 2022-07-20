package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please type a Hogwarts house name.")
		return
	}

	hogwarts := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dubmbledore", "lupin"},
		"hufflepuf": {"wenclock", "scamander", "helga", "digory"},
		"ravenclaw": {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin": {"horrace", "nigellus", "higgs", "scorpius"},
		"bobo": {"wizadry", "unwanted"},
	}

	delete(hogwarts, "bobo")
	key := args[1]

	if _, ok := hogwarts[key]; !ok {
		fmt.Printf("Sorry, I don't know anything about %s", key)
		return
	}

	names := hogwarts[key]
	sort.Strings(names)
	fmt.Printf("~~~ %s students ~~~\n\n", key)

	for _, name := range names {
		fmt.Println("\t+ " + name)
	} 
}