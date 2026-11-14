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
	var prioridad int = 3
	completada := false
	const estadoInicial string = "Pendiente"

	fmt.Println("---- INFORMACION DE LA TAREA ----")
	fmt.Println("Estado inicial:", estadoInicial)
	fmt.Println("Titulo:", titulo)
	fmt.Println("Descripcion:", descripcion)
	fmt.Println("Prioridad:", prioridad)
	fmt.Println("Completada:", completada)

	numero1 := 10
	numero2 := 3
	fmt.Println("---- OPERACIONES MATEMATICAS ----")
	fmt.Println("Suma:", numero1+numero2)
	fmt.Println("Resta:", numero1-numero2)
	fmt.Println("Multiplicacion:", numero1*numero2)
	fmt.Println("Division:", numero1/numero2)
	fmt.Println("Resto:", numero1%numero2)

	fmt.Println("-----OPERACIONES DE COMPRACION----")

	fmt.Println("¿Prioridad igual a 3?", prioridad == 3)
	fmt.Println("¿Prioridad diferente de 2?", prioridad != 2)
	fmt.Println("¿Prioridad mayor que 1?", prioridad > 1)
	fmt.Println("¿Prioridad menor que 5?", prioridad < 5)
	fmt.Println("¿Prioridad mayor o igual a 3?", prioridad >= 3)
	fmt.Println("¿Prioridad menor o igual a 2?", prioridad <= 2)
}
