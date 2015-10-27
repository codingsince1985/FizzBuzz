package main

import . "fmt"

func main() {
	for i := 1; i < 101; i++ {
		a := 0
		if i%3 == 0 {
			a, _ = Print("Fizz")
		}
		if i%5 == 0 {
			if a > 0 {
				Print("-")
			}
			a, _ = Print("Buzz")
		}
		if a < 1 {
			Print(i)
		}
		Println()
	}
}
