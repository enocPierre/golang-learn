package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14
	var raio = 3.2 // tipo float inferido pelo compilador

	// forma reduzinda de criar var
	area := PI * math.Pow(raio, 2)

	fmt.Print(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)
}
