package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fizzOrBuzz := false

		if i%3 == 0 {
			fizzOrBuzz = true
			fmt.Print("Fizz")
		}

		if i%5 == 0 {
			fizzOrBuzz = true
			fmt.Print("Buzz")
		}

		if !fizzOrBuzz {
			fmt.Print(i)
		}

		fmt.Println()
	}
}
