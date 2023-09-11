package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")

	// constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// variables
	alto := 15 // el dos puntos se usa cuando la variable se crea por primera vez
	var ancho int = 10
	var largo int // zero value

	fmt.Println(alto, ancho, largo)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a) // 0
	fmt.Println(b) // 0
	fmt.Println(c) // ""
	fmt.Println(d) // false

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area Cuadrado =", areaCuadrado)
}