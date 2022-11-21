package paquete

import "fmt"

// 'salario' solo se puede ver
// en el paquete 'paquete'
var salario = 12345.67

// Edad la puede ver el paquete
// y aquellos que importen el paquete
// Al poner la primera letra en
// mayusculas, se convierte en publica
var Edad = 47

// funcionPrivada solo es visible
// dentro del paquete 'paquete'
func funcionPrivada() {
	// 'esCasado' solo es visible en este metodo
	var esCasado = true

	fmt.Println("paquete > funcionPrivada() > salario: ", salario)
	fmt.Println("paquete > funcionPrivada() > Edad: ", Edad)
	fmt.Println("paquete > funcionPrivada() > esCasado: ", esCasado)
}

// FuncionPublica es visible
// dentro del paquete y donde se importe el paquete
// Al poner la primera letra en mayusculas,
// se convierte en metodo publico
func FuncionPublica() {
	// 'sabeGo' solo es visible en este metodo
	var sabeGo = true

	fmt.Println("paquete > FuncionPublica() > salario: ", salario)
	fmt.Println("paquete > FuncionPublica() > Edad: ", Edad)
	fmt.Println("paquete > FuncionPublica() > sabeGo: ", sabeGo)

	// Invoca a funcion privada en mismo archivo y paquete
	funcionPrivada()
}

// Calculo es una funcion publica
// Invoca a una funcion privada en otro archivo
// del mismo paquete, y retorna el resultado
func Calculo(a, b int) int {
	return sumatorio(a, b)
}
