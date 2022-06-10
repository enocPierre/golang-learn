package main

import (
	"fmt"
)

func main() {
	//var aporavado map[int] string
	//mapas devem ser inicializadosconst

	approvados := make(map[int]string)

	approvados[1234567] = "Maria"
	approvados[1234568] = "Pedro"
	approvados[1234569] = "Ana"

	fmt.Println(approvados)


	for cpf, nome := range approvados {
		fmt.Printf("%s (cpf: %d)\n", nome, cpf)
	}

	fmt.Println(approvados[1234567])
	delete(approvados, 1234567)
}
