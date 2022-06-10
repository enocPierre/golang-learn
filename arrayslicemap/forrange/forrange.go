package main

import "fmt"


func main() {
	numeros := [...] int {1,2,3,4,5} // compilador contas!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for num := range numeros {
		fmt.Println(num)
	}
	for _, num := range numeros {
		fmt.Println(num)
	}
}