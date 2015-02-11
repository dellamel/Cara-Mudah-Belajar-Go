package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 67
	x[1] = 78
	x[2] = 84
	x[3] = 98
	x[4] = 79
	var total float64 = 0.0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
