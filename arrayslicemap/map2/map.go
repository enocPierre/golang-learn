package main

import (
	"fmt"
)

func main() {
	funccsSalarios := map[string] float64 {
		"José jao" : 11325.45,
		"Gabriel Junior" : 1200.0,
		"Pedro Junior" : 15456.78,
	}

	funccsSalarios["Rafael Filho"] = 1350.0
	delete(funccsSalarios, "inexistente")


	for nome, salario := range funccsSalarios {
		fmt.Println(nome, salario)
	}
}