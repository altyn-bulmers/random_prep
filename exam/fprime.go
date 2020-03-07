package main

import "fmt"

func main() {
	primes(225225)
}

// func primes(n int) {
// 	for n%2 == 0 {
// 		fmt.Print("2 * ")
// 		n /= 2
// 	}
// 	for i := 3; i*i <= n; i += 2 {
// 		if n%i == 0 {
// 			fmt.Printf("%d * ", i)
// 			n /= i
// 		}
// 	}
// 	if isPrime(n) && n > 2 {
// 		fmt.Print(n, "\n")
// 	}
// }

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= n/3; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func primes(n int) {
	result := ""
	for n%2 == 0 {
		result = result + "2"
		n /= 2
	}
	//	index := 0
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			result += string(i)
			n /= i
			i = 3
			//			index = i
		}
	}
	if isPrime(n) && n > 2 {
		result += string(n)
	}
	for i := 0; i < len(result)-1; i++ {
		fmt.Print(result[i], " * ")
	}
	fmt.Println(result[len(result)-1])
}
