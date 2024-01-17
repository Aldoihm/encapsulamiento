package main

import (
	"fmt"

	"github.com/Aldoihm/encapsulamiento/course"
)

func main() {
	Go := course.New("Go desde Cero", 0, true)
	Go.SetUserIDs([]uint{7, 60, 1, 20})
	Go.SetClasses(
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		})

	fmt.Printf("%+v \n", Go)
	Go.SetName("POO en Go")
	Go.SetPrice(180)
	Go.SetIsFree(false)

	fmt.Println("Name: ", Go.Name())
	fmt.Println("Price: ", Go.Price())
	fmt.Println("Is Free?: ", Go.IsFree())

}
