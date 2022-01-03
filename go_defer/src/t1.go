package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("step 1")
	}()

	defer func() {
		fmt.Println("step 2")
	}()

	defer func() {
		fmt.Println("step 3")
	}()

	fmt.Println("step 4")
}
