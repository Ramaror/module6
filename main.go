package main

import "fmt"

type cat struct {
	a, b int
}
type dog struct {
}

type Animal interface {
	Run() int
	Sound() int
}

func (c *cat) Run() int {
	a := c.a + c.b
	return a
}

func (d *cat) Sound() int {
	b := d.a * d.b
	return b
}
func main() {
	var a Animal
	sh := cat{1, 2}
	a = &sh
	fmt.Println(a.Sound())
	fmt.Println(a.Run())
}
