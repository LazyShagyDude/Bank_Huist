package main

import "fmt"

func main() {
	for i := 1; i <= 200; i++ { // 200 numbers
		found := false

		if i%3 == 0 {
			fmt.Printf("Fizz")
			found = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			found = true
		}

		if !found {
			fmt.Println(i)
			continue
		}

		fmt.Println()
	}
}
