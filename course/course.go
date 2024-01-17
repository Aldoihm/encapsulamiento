package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// c Course.. es una copia del valor
func (c Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// Como quiero hacer un cambio, al valor original y no a la copia, pongo de receptor un puntero
func (c *Course) changePrice(price float64) {
	c.Price = price
}
