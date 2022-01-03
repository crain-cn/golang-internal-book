package main

import "fmt"

func main() {
	for i := 1; i < 4; i++   {
		defer func() {
			fmt.Println("step ",i)
		}()
	}
}
