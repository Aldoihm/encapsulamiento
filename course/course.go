package course

import "fmt"

type Course struct {
	Name     string
	Price    float64
	IsFree   bool
	UsersIDs []uint
	Classes  map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c *Course) PrintClases() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
