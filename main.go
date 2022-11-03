package main

import "github.com/Aldoihm/go-poo/tree/master/encapsulation/course"

func main() {
	goland := course.Course{
		Name:  "Go",
		Price: 12.34,
		Classes: map[uint]string{
			1: "Variables",
			2: "Tipos de Datos",
			3: "Operadores",
			4: "Condicionales",
		},
	}
	//Así llamamos un método en Go, igual es diferente a las funciones
	goland.PrintClases()
}
