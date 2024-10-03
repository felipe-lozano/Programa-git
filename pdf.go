package main

import (
	"log"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// Crear un nuevo PDF
	pdf := gofpdf.New("P", "mm", "A4", "") // Inicializa el PDF
	pdf.AddPage()

	// Establecer fuente
	pdf.SetFont("Arial", "", 14)

	// Escribir texto
	pdf.Cell(0, 10, "Hola, este es un PDF simple!") // Ancho 0 para que ocupe toda la línea
	pdf.Ln(10) // Salto de línea
	pdf.Cell(0, 10, "Generado usando la biblioteca gofpdf en Go.")

	// Guardar el PDF
	err := pdf.OutputFileAndClose("simple.pdf")
	if err != nil {
		log.Fatal(err)
	}
}