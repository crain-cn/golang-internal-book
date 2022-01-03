package main

import "fmt"

func main() {
	i := 0
	defer func() {
		fmt.Println("step i ",i)
	}()
	i++

	j := 0
	defer func(j int) {
		fmt.Println("step  j ",j)
	}(j + 1)
	j++

	fmt.Println("step  end ")
}
