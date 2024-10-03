package main

import (
	"fmt"
	
)

func main() {
	var descripcion string
	var cantidadStr int
	var precioStr int
	var totalFinal int

	// Pedir al usuario la cantidad de productos
	fmt.Print("¿Cuántos productos desea agregar? ")
	var numProductos int
	fmt.Scanln(&numProductos)

	for i := 0; i < numProductos; i++ {


		// Leer descripción del producto
		fmt.Print("producto " ,i+1, ": ")
		fmt.Scanln(&descripcion)

		// Leer cantidad
		fmt.Print("Cantidad: ")
		fmt.Scanln(&cantidadStr)
		cantidad := cantidadStr

		// Leer precio unitario
		fmt.Print("Precio unitario: ")
		fmt.Scanln(&precioStr)
		precioUnitario := precioStr

		// Calcular total para este producto
		total := cantidad * precioUnitario
		totalFinal += total

		// Mostrar el producto en consola
		fmt.Println("producto: ",descripcion, " cantidad: ",cantidad, " precio unitario: ",precioUnitario, " total: ", total )
	}

	// Mostrar el total final de la factura
	fmt.Printf("\nTotal Final de la Factura: %d\n", totalFinal)
}

