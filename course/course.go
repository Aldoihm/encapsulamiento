package course

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// Función constructora
func New(name string, price float64, isfree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isfree,
	}
}

// c Course.. es una copia del valor
func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// Como quiero hacer un cambio, al valor original y no a la copia, pongo de receptor un puntero
func (c *course) changePrice(price float64) {
	c.Price = price
}

//funcion constructora
