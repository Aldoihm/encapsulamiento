package course

import "fmt"

type course struct {
	Name     string
	Price    float64
	IsFree   bool
	UsersIDs []uint
	Classes  map[uint]string
}

func (c *course) ChangePrice(price float64) {
	c.Price = price
}

func (c *course) PrintClases() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

//Funcion constructora
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}
