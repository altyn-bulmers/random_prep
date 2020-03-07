package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args
	numb, _ := strconv.Atoi(arg[1])
	Fprime(numb)
}

func Fprime(n int) {
	a := 2
	for a <= n {
		if n%a == 0 {
			n = n / a
			if a <= n {
				fmt.Print(a, " * ")
			} else {
				fmt.Print(a, "\n")
			}
		} else {
			a++
		}
	}
}
