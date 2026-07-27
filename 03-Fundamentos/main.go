package main

import "fmt"

const tax = 0.16

func main() {
	var (
		name     string
		email    string
		subtotal float64
	)

	fmt.Println("Ticket por consola")

	fmt.Println("Nombre del cliente: ")
	fmt.Scan(&name)

	fmt.Println("Email del cliente: ")
	fmt.Scan(&email)

	fmt.Println("Subtotal de productos:")
	fmt.Scan(&subtotal)

	fmt.Println("**************************************************")
	fmt.Printf("Subtotal:	        %.2f\n", subtotal)
	fmt.Printf("Impuestos (IVA):	%.2f\n", subtotal*tax)
	fmt.Printf("Total:				%.2f\n", subtotal+(subtotal*tax))
	fmt.Println("**************************************************")

	// Ticket con impuestos
	// Nombre del  cliente
	// Email del cliente
	// Subtotal

	//Respuest:
	// Subtotal
	// Impuestos (IVA)
	// Total (Ya con impuestos incluidos)
}
