package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. Define el struct Tarea aquí
type Tarea struct {
	ID          int
	Descripcion string
	Completada  bool
}

// 2. Variable global para el slice de tareas y el contador de ID
var tareas = []Tarea{}
var nextID = 1

// Implementa estas funciones:
func agregarTarea(scanner *bufio.Scanner) {
	fmt.Print("Ingresa la descripción de la tarea: ")
	scanner.Scan()
	descripcion := strings.TrimSpace(scanner.Text())

	tarea := Tarea{
		ID:          nextID,
		Descripcion: descripcion,
		Completada:  false,
	}
	tareas = append(tareas, tarea)
	nextID++

	fmt.Println("Tarea agregada con éxito.")
}

func listarTareas() {
	if len(tareas) == 0 {
		fmt.Println("No hay tareas pendientes.")
		return
	}

	fmt.Println("\n=== LISTA DE TAREAS ===")
	for _, tarea := range tareas {
		status := "Pendiente"
		if tarea.Completada {
			status = "Completada"
		}
		fmt.Printf("ID: %d | Descripción: %s | Estado: %s\n", tarea.ID, tarea.Descripcion, status)
	}
}

func completarTarea(scanner *bufio.Scanner) {
	fmt.Print("Ingresa el ID de la tarea a completar: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID inválido. Intenta de nuevo.")
		return
	}

	for i := range tareas {
		if tareas[i].ID == id {
			tareas[i].Completada = true
			fmt.Println("Tarea marcada como completada.")
			return
		}
	}

	fmt.Printf("No se encontró una tarea con ese ID %d.\n", id)
}

func eliminarTarea(scanner *bufio.Scanner) {
	fmt.Print("Ingresa el ID de la tarea a eliminar: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("ID inválido. Intenta de nuevo.")
		return
	}

	for i, tarea := range tareas {
		if tarea.ID == id {
			tareas = append(tareas[:i], tareas[i+1:]...)
			fmt.Println("Tarea eliminada con éxito.")
			return
		}
	}

	fmt.Printf("No se encontró una tarea con ese ID %d.\n", id)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== GESTOR DE TAREAS ===")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Listar tareas")
		fmt.Println("3. Completar tarea")
		fmt.Println("4. Eliminar tarea")
		fmt.Println("5. Salir")
		fmt.Print("Elige una opción: ")

		scanner.Scan()
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			agregarTarea(scanner)
		case "2":
			listarTareas()
		case "3":
			completarTarea(scanner)
		case "4":
			eliminarTarea(scanner)
		case "5":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida. Intenta de nuevo.")
		}
	}
}
