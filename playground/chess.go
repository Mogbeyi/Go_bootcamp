package main

import "fmt"

type File []bool

type Chessboard map[string]File

func main() {
	cb := newChessboard()
	
	for _, sq := range cb {
		fmt.Println(sq)
	}
}

func newChessboard() Chessboard {
	return Chessboard{ 
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, true, false, false, false, false, false},
		"D": File{false, false, false, false, false, false, false, false},
		"E": File{false, false, false, false, false, true, false, true},
		"F": File{false, false, false, false, false, false, false, false},
		"G": File{false, false, false, true, false, false, false, false},
		"H": File{true, true, true, true, true, true, false, true},
	}
}


