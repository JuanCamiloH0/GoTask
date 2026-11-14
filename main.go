package main

import "fmt"

const Project = "GoTask"
const Language = "Go"
const Version = "1.0"

func main() {
	fmt.Println("================================")
	fmt.Println("          GoTask v1.0")
	fmt.Println("================================")
	fmt.Println("")
	fmt.Println("Administrador de tareas")
	fmt.Println("")
	fmt.Println("Proyecto:", Project)
	fmt.Println("Lenguaje:", Language)
	fmt.Println("Version:", Version)
	fmt.Println("")
	fmt.Println("¡Bienvenido!")

	var titulo string = "Aprender Go"
	descripcion := "Estudiar variables y tipos de datos en Go"
	var prioridad int = 2
	completado := false

	fmt.Println("---- INFORMACION DE LA TAREA ----")
	fmt.Println("Titulo:", titulo)
	fmt.Println("Descripcion:", descripcion)
	fmt.Println("Prioridad:", prioridad)
	fmt.Println("Completado:", completado)
}
