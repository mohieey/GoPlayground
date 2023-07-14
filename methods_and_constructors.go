package main

type Cat struct {
	name  string
	age   int
	color string
}

func (c *Cat) talk() {
	println("Meawoooooooooooooooo")
}

func newCat() *Cat {
	return &Cat{
		color: "black",
	}
}
