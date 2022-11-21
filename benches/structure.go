package main

import (
	"fmt"
)

type User struct {
	name        string
	age         int
	bankBalance float32
}

func main() {
	var user User
	fmt.Println(user)
}
