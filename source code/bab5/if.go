package main

import "fmt"

func main() {
	for a := 1; a <= 9; a++ {
		if a%2 == 0 {
			fmt.Println(a, "genap")
		} else {
			fmt.Println(a, "ganjil")
		}
	}
}
