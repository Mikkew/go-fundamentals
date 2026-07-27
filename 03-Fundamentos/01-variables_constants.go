package main

import "fmt"

const AppName = "Curso Go 2.0" // app = Interno, App = Exportable
const MaxUsers = 1000

func constants() {
	var name string = "Miguel"
	lastName := "Medina"

	fmt.Printf("Nombre: %s %s %T \n", name, lastName, lastName)
	fmt.Println(AppName, MaxUsers)
}
