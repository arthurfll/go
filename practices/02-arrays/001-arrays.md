```go




package main

import ("fmt"; "reflect")


func main() {

	lista := [5]int {1,2,3,4}
	var lista2 [5]int

	fmt.Println(reflect.TypeOf(lista))
	fmt.Println(reflect.TypeOf(lista2))

}
