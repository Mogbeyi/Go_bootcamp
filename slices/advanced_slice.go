package main

import (
	// "fmt"
	s "github.com/inancgumus/prettyslice"
)

func main() {
	names := make([]string, 5)
	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	s.Show("1st step", names)
	
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("2nd step", names)

	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("3rd step", names)

	copy(names[3:5], moreNames[:2])
	names = names[:cap(names)]
	s.Show("4th step", names)

	clones := make([]string, 3, 5)
	copy(clones, names[2:])
	s.Show("5th step", clones)
	clones = append(clones, names[:2]...)
	s.Show("5th step", clones)

	sliced := clones[1:5]
	sliced = append(sliced, "hypatia")
	s.Show("6th step", sliced)
	clones[2] = "elder"
	s.Show("6th step", clones, sliced)
}
