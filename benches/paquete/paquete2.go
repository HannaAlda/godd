package paquete

// PI es una constante publica
const PI = 3.141519

// AreaCircunferencia calcula el area de
// la circunferencia a partir de su radio
// Funcion publica
func AreaCircunferencia(radio float64) float64 {
	return PI * (radio * radio)
}

// Funcion privada. Solo es visible en el paquete
func sumatorio(a, b int) int {
	return a + b
}
