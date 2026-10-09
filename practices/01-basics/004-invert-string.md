```go




package main

import "fmt"

func main() {
	x := "arthur"

    y := len(x)

	saida := ""

	for i:= y-1; i>=0 ; i-- {
		saida = saida + string(x[i])

	}
	fmt.Println(saida)
}

