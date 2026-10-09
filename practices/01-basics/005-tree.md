```go



package main

import "fmt"

func main() {
    
	output := ""

	for i:= 0 ; i<5; i++ {

		for j := 0; j<=i ; j++ {

			output = output + "*"
		}

		output = output+"\n"
	}
	fmt.Println(output)
}


