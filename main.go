package main

import("fmt"; "reflect"; "strconv")

func main() {
	var x int
	var y int
	z := x + y

	z1 := strconv.Itoa(z)

	fmt.Println(reflect.TypeOf(z))
	fmt.Println(reflect.TypeOf(z1))

	var z2 string
	z2 = fmt.Sprint("número "+z1)

	fmt.Println(z2)
}


