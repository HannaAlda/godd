package main

import "fmt"

type Thing struct {
	Name string
	Num  int
}

func (t *Thing) Init(name string, num int) {
	t.Name = name
	t.Num = num
}

func main() {
	t := new(Thing)
	t.Init("Hello Boss!", 99)
	fmt.Printf("%s: %d\n", t.Name, t.Num)
}
