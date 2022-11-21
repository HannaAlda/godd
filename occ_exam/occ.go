package main

import "fmt"

func main() {
	str := "Banana"
	strRev := "Impar"
	if par(str) {
		strRev = reverse(str)
	}
	fmt.Println(str)
	fmt.Println(strRev)

}

func par(s string) bool {
	length := len(s)
	if length%2 == 0 {
		return true
	}

	return false
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
