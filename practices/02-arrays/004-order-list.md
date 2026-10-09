```go


package main

import "fmt"

func main() {
	lista := []int{7,3,9,6,1,8,2,5,0,4}
	var temp int
	var cont int

	for j:=0 ; j < len(lista)-1 ; j++ {
		cont = 0
		for i:=0 ; i < len(lista)-1 ; i++ {

			if lista[i] > lista[i+1] {
				temp = lista[i]
				lista[i] = lista[i+1]
				lista[i+1] = temp
				cont++
			}
		if cont == 0 {
			break
		}
	}
	}

	fmt.Println(lista)
}


