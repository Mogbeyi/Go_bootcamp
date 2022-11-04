package main

import "fmt"

type error string

func (e error) String() string {
	return fmt.Sprintf("MyErr: %v", string(e))
}
