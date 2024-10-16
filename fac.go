package main

import (
	"fmt"
	"time"
	
)
func main() {
	// Obtener la fecha y hora actual
	currentTime := time.Now()

	// Formatear la fecha en formato deseado (día/mes/año)
	fmt.Println("Fecha actual:", currentTime.Format("02-01-2006"))
}
