package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Somar =>", a+b)
	fmt.Println("Somar =>", a-b)
	fmt.Println("Somar =>", a/b)
	fmt.Println("Somar =>", a*b)
	fmt.Println("Somar =>", a%b)

	//betwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("xor =>", a^b)

	c := 3.0
	d := 2.0

	// outras operacoes usando math..
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciacao =>", math.Pow(c, d))



}