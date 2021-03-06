package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//numeros inteiros
	fmt.Println(1, 2, 1000)
	//reflect dizer tipo da variavel quando passa o typeOf
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos) .. unit8, unit16, unit32, unit64
	var b byte = 225
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal .. int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x = float32(49.99)
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 ", reflect.TypeOf(49.99))//float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é, ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// Sting 
	s1 := "Ola meu nome é Pierre"
	fmt.Println("O tamanha da string é ", len(s1))

	// char ??
	char := 'a'
	fmt.Println("O tipo de char é: ", reflect.TypeOf(char))

}
