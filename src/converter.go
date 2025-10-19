package main

import "fmt"

func main() {
	for {
		choice := showMenuAndGetChoice()

		switch choice {
		case 1:
			fmt.Print("Digite a temperatura em Celsius: ")
			var tempC float64
			fmt.Scanln(&tempC)
			fmt.Printf("\nResultado: %.2f°C é igual a %.2f°K\n\n", tempC, CelsiusToKelvin(tempC))
		case 2:
			fmt.Print("Digite a temperatura em Kelvin: ")
			var tempK float64
			fmt.Scanln(&tempK)
			fmt.Printf("\nResultado: %.2fK é igual a %.2f°C\n\n", tempK, KelvinToCelsius(tempK))
		case 3:
			fmt.Println("Saindo do programa. Até logo!")
			return
		default:
			fmt.Printf("\nOpção %d inválida! Por favor, tente novamente.\n\n", choice)
		}
	}
}

func showMenuAndGetChoice() int {
	for {
		fmt.Println("##############################################################################################")
		fmt.Println("################################## Conversor de Temperatura ##################################")
		fmt.Println("##############################################################################################")
		fmt.Println("Escolha uma opção de conversão:")
		fmt.Println("1: Celsius para Kelvin")
		fmt.Println("2: Kelvin para Celsius")
		fmt.Println("3: Sair")
		fmt.Println("##############################################################################################")
		fmt.Print("Digite sua escolha: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número (1, 2 ou 3).")
			// Limpa o buffer de entrada para evitar um loop infinito se o usuário digitar texto
			var discard string
			fmt.Scanln(&discard)
			continue // Pede a entrada novamente
		}
		return choice // Retorna a escolha se for um número válido
	}
}

func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}
