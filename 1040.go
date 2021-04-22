package main

import (
	"fmt"
)

func main() {
	var n [4]float64
	var exame float64
	var soma float64 = 0.0

	for i := 0; i < 4; i++ {
		fmt.Scan(&n[i])
	}
	soma = (n[0] * 2) + (n[1] * 3) + (n[2] * 4) + n[3]
	soma = soma / 10
	s := fmt.Sprintf("Media: %.1f\n", soma) //(math.Floor(soma*10) / 10))
	fmt.Print(s)
	if soma >= 7 {
		fmt.Println("Aluno aprovado.")
	} else if soma >= 5 {
		fmt.Println("Aluno em exame.")
		fmt.Scan(&exame)
		s := fmt.Sprintf("Nota do exame: %.1f\n", exame)
		fmt.Print(s)
		if exame >= 5 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		soma = (soma + exame) / 2
		t := fmt.Sprintf("Media final: %.1f\n", soma)
		fmt.Print(t)
	} else {
		fmt.Println("Aluno reprovado.")
	}
}
