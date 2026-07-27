package main

import "fmt"

//Crear dos structs

//Item -> SKU, Name, Price, Qty
//Order -> ID, Customer, Items, Meta (City, Source)

/*
	1. Asignar valores al order
	2. Imprimir
	   - Order ID
	   - Cliente
	   - Acceso al primer elemento del item
	   - Acceso directo a un valor en el map de Meta.
	   - Modificar el cupon a HOLACARADEBOLA
	   - Mostrar el cupon modificado
*/

type Item struct {
	SKU   string
	Name  string
	Price float64
	Qty   int
}

type Order struct {
	ID       string
	Customer string
	Items    []Item
	Meta     map[string]string
}

func main() {
	order := Order{
		ID:       "ORDER-123",
		Customer: "Alfredo Samano",
		Items: []Item{
			{SKU: "ZPTS-0234", Name: "Zapatos", Price: 120.50, Qty: 2},
			{SKU: "CAMI-5678", Name: "Camisa", Price: 45.99, Qty: 1},
			{SKU: "PANT-9012", Name: "Pantalón", Price: 60.00, Qty: 3},
			{SKU: "GAF-3456", Name: "Gafas de sol", Price: 80.00, Qty: 1},
		},
		Meta: map[string]string{
			"City":   "Ciudad de México",
			"Source": "Online",
		},
	}

	fmt.Println("Order ID: ", order.ID)
	fmt.Println("Cliente: ", order.Customer)
	fmt.Println("Primer Item: ", order.Items[0].Name, ", SKU: ", order.Items[0].SKU, ", Precio: ", order.Items[0].Price, ", Cantidad: ", order.Items[0].Qty)
	fmt.Println("Ciudad: ", order.Meta["City"])
	order.Meta["Coupon"] = "MUNDIAL2026MEX"
	fmt.Println("Coupon: ", order.Meta["Coupon"])
	fmt.Println(order)
}
