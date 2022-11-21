package main

import (
	"fmt"
	"paquete"
)

func main() {
	// Acceso a las variables publicas
	fmt.Println("main() > paquete.Edad: ", paquete.Edad)
	fmt.Println("main() > paquete.PI: ", paquete.PI)

	// Invocaciones de funciones publicas
	paquete.FuncionPublica()

	var resultado = paquete.AreaCircunferencia(2.15)
	fmt.Println("main() > resultado: ", resultado)

	var suma int = paquete.Calculo(12, 5)
	fmt.Println("main() > suma: ", suma)
}
