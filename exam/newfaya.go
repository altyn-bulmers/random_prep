package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	first := args[0]
	second := args[1]

	result := ""
	index := 0

	for _, s := range first {
		for j := index; j < len(second); j++ {
			if byte(s) == second[j] {
				result = result + string(s)
				index = j + 1
				break
			}

		}
	}
	fmt.Println(makeUnique(result))
}

// func makeUnique(str string) string {
// 	result := ""

// 	for i := 0; i < len(str); i++ {
// 		exists := false

// 		for j := 0; j < i; j++ {

// 			if str[i] == str[j] {
// 				exists = true
// 				break
// 			}
// 		}
// 		if !exists {
// 			result = result + string(str[i])
// 		}
// 	}
// 	return result

// }

func makeUnique(str string) string {
	result := ""

	for i := 0; i < len(str); i++ {
		exists := false
		for j := 0; j < i; j++ {

			if str[i] == str[j] {
				exists = true
				break
			}
		}
		if !exists {
			result = result + string(str[i])
		}

	}
	return result
}
