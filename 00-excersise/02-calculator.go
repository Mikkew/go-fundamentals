package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar() {
	a := leerNumero("Ingresa el primer número: ")
	b := leerNumero("Ingresa el segundo número: ")

	resultado := a + b
	fmt.Printf("Resultado: %.2f + %.2f = %.2f\n", a, b, resultado)
}

func restar() {
	a := leerNumero("Ingresa el primer número: ")
	b := leerNumero("Ingresa el segundo número: ")

	resultado := a - b
	fmt.Printf("Resultado: %.2f - %.2f = %.2f\n", a, b, resultado)
}

func multiplicar() {
	a := leerNumero("Ingresa el primer número: ")
	b := leerNumero("Ingresa el segundo número: ")

	resultado := a * b
	fmt.Printf("Resultado: %.2f * %.2f = %.2f\n", a, b, resultado)
}

func dividir() {
	a := leerNumero("Ingresa el primer número: ")
	b := leerNumero("Ingresa el segundo número: ")

	if b == 0 {
		fmt.Println("❌ Error: No se puede dividir entre cero.")
		return
	}

	resultado := a / b
	fmt.Printf("Resultado: %.2f / %.2f = %.2f\n", a, b, resultado)
}

// Función auxiliar para leer un número (con validación)
func leerNumero(mensaje string) float64 {
	var (
		input  string
		numero float64
		err    error
	)

	for {
		fmt.Print(mensaje)
		fmt.Scanln(&input)

		numero, err = strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("❌ Entrada inválida. Por favor, ingresa un número válido.")
			continue
		}
		break
	}
	return numero
}

func main() {
	// Variable para controlar el bucle
	var continuar bool = true

	for continuar {
		// Mostrar menú
		fmt.Println("\n=== CALCULADORA ===")
		fmt.Println("1. Sumar")
		fmt.Println("2. Restar")
		fmt.Println("3. Multiplicar")
		fmt.Println("4. Dividir")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción (1-5): ")

		// Leer opción
		var opcion int
		fmt.Scanln(&opcion)

		// Procesar opción
		switch opcion {
		case 1:
			sumar()
		case 2:
			restar()
		case 3:
			multiplicar()
		case 4:
			dividir()
		case 5:
			fmt.Println("¡Hasta luego!")
			continuar = false
		default:
			fmt.Println("❌ Opción no válida. Intenta de nuevo.")
		}
	}
}
