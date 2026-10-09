```go


package main

import "fmt"

func main() {
	nota1 := 3.0
	nota2 := 2.5
	frequencia := 50
	trabalhos := 3

	media := (nota1 + nota2)/2

	if media >= 7.0 && frequencia >= 75 {
		fmt.Println("Aluno aprovado")
	} else if media >= 5.0 || trabalhos >= 8 {
		fmt.Println("Aluno de recuperação")
	} else {
		fmt.Println("Aluno reprovado")
	}
}

