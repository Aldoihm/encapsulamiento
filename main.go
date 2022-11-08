package main

import (
	"fmt"

	"github.com/Aldoihm/encapsulamiento/course"
)

func main() {
	goland := course.New("go", 1, false)
	goland.SetName("Poo con Go")
	goland.SetPrice(5.50)
	goland.SetIsFree(true)
	goland.SetUsersIDs([]uint{12, 56, 89})
	goland.SetClases(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	fmt.Println(goland.Name())
	fmt.Println(goland.Price())
	fmt.Println(goland.IsFree())
	fmt.Println(goland.UsersIDs())
	fmt.Println(goland.Classes())

	goland.PrintClases()

	/*
		goland.UsersIDs = []uint{12, 56, 89}
		goland.Classes = map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		} */
	//fmt.Println(goland.Classes)

}

//En Go no es necesario los métodos setters y getters
//En este ejercicio, cambiamos el campo de Classes como no exportados y lo dejamos en classes

//PASOS DE COMO ENCAPSULE ESTE PROYECTO
/*
	1.-subi a git mi proyecto
	2.-Estando en la carpeta encapsulimiento (raiz, aqui es donde se ejecuta go mod init) ejecute el comando:
		go mod init github.com/Aldoihm/encapsulamiento
		*con esta dirección después podré ejecutar el comando go get para obtener mis paquetes que yo hago
		*En este ejemplo tenia mi paquete en github.com/Aldoihm/encapsulamiento/course
		*fijarme que no ejecute el comando con la dirección: "github.com/Aldoihm/encapsulamiento/course"
		esto es un error, por eso puse hasta encapsulamiento, para que después con go get obtenga todos los paquetes
		que hice en las subdireccion
	3.-Ejecute el comando go get:
		go get github.com/Aldoihm/encapsulamiento/course
		*Esto porque esa fue la dirección que escribí en el import en el archivo main
		*No copiar la direccion que me daba github que era la siguiente:
			https://github.com/Aldoihm/encapsulamiento/tree/main/course
		Si se llega a copiar, empiezan a salir el montonal de errores
	4.- go build
	5.- Listo

	conclusión:
	1.- go mod init github.com/Aldoihm/encapsulamiento
	2.- go get github.com/Aldoihm/encapsulamiento/course
	3.- go build
*/
