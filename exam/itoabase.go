package main

import "fmt"

func main() {
	fmt.Println(itoaBase(-125, 4))
}

func itoaBase(value, base int) string {
	origin := "0123456789ABCDEF"
	str := origin
	str = str[:base]
	runes := []rune(str)
	result := ""
	if base < 2 {
		return "NV"
	}
	if value == 0 {
		return "0"
	}
	if value < 0 {
		result += "-"
	}
	result = result + baseRecursion(value, base, runes)
	return result
}

func baseRecursion(nbr int, base int, runes []rune) string {
	result := ""
	if nbr/base != 0 {
		result += baseRecursion(nbr/base, base, runes)
	}

	mod := 0
	mod = nbr % base
	if mod < 0 {
		mod = -mod
	}
	result = result + string(runes[mod])
	return result
}
