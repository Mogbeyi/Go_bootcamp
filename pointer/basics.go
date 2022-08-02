package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	//create a null pointer with the type of pointer to a computer
	var null *computer

	//Compare the pointer variable to a nil value
	if null == nil {
		fmt.Println("null computer is nil")
	}

	//Create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}

	// put the apple into a new pointer variable
	newApple := apple

	// compare the apples: if they are equal
	if apple == newApple {
		// say so and print their addresses
		fmt.Printf("apples are equal          : apple: %p newApple: %p\n",
			apple, newApple)
	}

	sony := &computer{brand: "sony"}

	if apple != sony {
		// say so and print their addresses
		fmt.Printf("apple and sony are inequal: apple: %p sony: %p\n",
			apple, sony)
	}

	appleVal := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple                     : %p %p\n", &apple, apple)
	fmt.Printf("appleVal                  : %p\n", &appleVal)

	if *apple == appleVal {
		fmt.Println("apple and appleVal are equal")

		// print the values:
		// the value that is pointed by the apple and the new variable
		fmt.Printf("apple                     : %+v — appleVal: %+v\n",
			*apple, appleVal)
	}

	// change sony's brand to hp using the func
	change(sony, "hp")
	// print sony's brand
	fmt.Printf("sony                      : %s\n", sony.brand)

	// print the returned value
	fmt.Printf("appleVal                  : %+v\n", valueOf(apple))

	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))
}

func change(c *computer, brand string) {
	c.brand = brand
}

func valueOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}

