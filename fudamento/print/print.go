package main

import "fmt"

func main() {
	fmt.Print("Messa")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)

	xs := fmt.Sprint(x)
	fmt.Println("Ovalor de x é" + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("o valor é %.2f.", x)
		
	a := 1
	b := 1.9999
	c := false
	d := "opa"
	
	fmt.Printf("\n%d %f %.2f %t %v ", a, b, c, d)
	fmt.Printf("\n%d %.2f %v %s ", a, b, c, d)

}