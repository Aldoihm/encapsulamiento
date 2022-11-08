package course

import "fmt"

type course struct {
	name     string
	price    float64
	isFree   bool
	usersIDs []uint
	classes  map[uint]string
}

func (c *course) PrintClases() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

//Métodos setters
func (c *course) SetName(name string)              { c.name = name }
func (c *course) SetPrice(price float64)           { c.price = price }
func (c *course) SetIsFree(isfree bool)            { c.isFree = isfree }
func (c *course) SetUsersIDs(usersids []uint)      { c.usersIDs = usersids }
func (c *course) SetClases(clases map[uint]string) { c.classes = clases }

//Métodos getters
func (c *course) Name() string             { return c.name }
func (c *course) Price() float64           { return c.price }
func (c *course) IsFree() bool             { return c.isFree }
func (c *course) UsersIDs() []uint         { return c.usersIDs }
func (c *course) Classes() map[uint]string { return c.classes }

//Funcion constructora
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}
