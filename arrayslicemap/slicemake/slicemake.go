package main

import "fmt"


func main() {
	// criar slice apartir da make, internamente ele ja criar array associado 
	s := make([] int, 10)
	s[9] = 12
	fmt.Println(s)
	//s []int
	//slice parti make dizendo tamnho capcidade interno do array que vai criar 
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1,2,3,4,5,6,7,8,9,0)
	fmt.Println(s, len(s), cap(s))
    //slice  quando atingir ao tamanho maximo, ele doublo  
	s = append(s, 1)
	fmt.Println(s,len(s), cap(s))
}