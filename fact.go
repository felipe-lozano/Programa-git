package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estructura para almacenar los ítems
type Item struct {
	Descripcion   string
	Cantidad      int
	PrecioUnitario float64
}

func main() {
	var items []Item

	// Entrada del usuario para agregar productos
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("¿Desea añadir un producto a la factura? (s/n)")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)

		if confirm == "n" {
			break
		}

		// Leer descripción del producto
		fmt.Print("Descripción del producto: ")
		descripcion, _ := reader.ReadString('\n')
		descripcion = strings.TrimSpace(descripcion)

		// Leer cantidad
		fmt.Print("Cantidad: ")
		cantidadStr, _ := reader.ReadString('\n')
		cantidadStr = strings.TrimSpace(cantidadStr)
		cantidad, _ := strconv.Atoi(cantidadStr)

		// Leer precio unitario
		fmt.Print("Precio unitario: ")
		precioStr, _ := reader.ReadString('\n')
		precioStr = strings.TrimSpace(precioStr)
		precioUnitario, _ := strconv.ParseFloat(precioStr, 64)

		// Agregar el producto a la lista de items
		item := Item{
			Descripcion:   descripcion,
			Cantidad:      cantidad,
			PrecioUnitario: precioUnitario,
		}
		items = append(items, item)
	}

	// Mostrar la factura en consola
	mostrarFactura(items)
}

func mostrarFactura(items []Item) {
	// Mostrar detalles de la factura
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Factura")
	fmt.Println("Número de Factura: 12345")
	fmt.Println("Fecha: 01/10/2024")
	fmt.Println("Cliente: Juan Pérez")
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%-20s %-10s %-15s %-15s\n", "Descripción", "Cantidad", "Precio Unitario", "Total")
	fmt.Println("-----------------------------------------------------")

	// Agregar los ítems ingresados por el usuario y calcular total
	var totalFinal float64 = 0
	for _, item := range items {
		total := float64(item.Cantidad) * item.PrecioUnitario
		totalFinal += total
		fmt.Printf("%-20s %-10d %-15.2f %-15.2f\n", item.Descripcion, item.Cantidad, item.PrecioUnitario, total)
	}

	// Total final de la factura
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%-20s %-10s %-15s %-15.2f\n", "Total Final", "", "", totalFinal)
	fmt.Println("-----------------------------------------------------")
}
