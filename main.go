package main

import (
	"fmt"
)

func main() {
	fmt.Println("Digite a altura:")
	var altura float64
	fmt.Scanf("%f", &altura)
	fmt.Println("Digite o Peso")
	var peso float64
	fmt.Scanf("%f", &peso)
	var imc = peso / (altura * altura)
	fmt.Println("O seu IMC é de:", imc)

	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	} else if imc < 24.9 {
		fmt.Println("Prdo normal")
	} else if imc < 29.9 {
		fmt.Println("Acima do Peso")
	} else {
		fmt.Println("Obesidade")
	}
}
