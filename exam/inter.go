package main

import (
	"fmt"
)

func main() {
	str := "padinton"
	fmt.Println(makeUnique(str))
}

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
