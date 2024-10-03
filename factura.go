package main

import (
	"fmt"
)

func factura() {
	var producto string
	var precio int
	var cantidad int
	var valorfinal int

	fmt.Println("Ingrese el nombre del producto: ")
	fmt.Scanln(&producto)
	fmt.Println("Ingrese el precio del producto: ")
	fmt.Scanln(&precio)
	fmt.Println("Ingrese la cantidad del producto: ")
	fmt.Scanln(&cantidad)
	fmt.Println("quiere añadir otro producto: ")
	var respuesta string
	fmt.Scanln(&respuesta)
	if respuesta == "si" {
		factura()
	}

	valor := cantidad * precio
	valorfinal += valor
	fmt.Println(" Producto: ", producto, "precio", precio, "valor ", valor)
	fmt.Println("el valor final es ", valorfinal)

}

func main() {
	factura()
}
