package main

import (
	"fmt"

	"github.com/Aldoihm/encapsulamiento/course"
)

func main() {
	Go := course.New("Go desde Cero", 0, false)
	Go.UserIDs = []uint{12, 56, 98}
	Go.Classes = map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	}

	fmt.Printf("%+v", Go)
}
