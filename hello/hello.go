package grettings

import "fmt"

func Hola(nombre string) string {
	message := fmt.Sprintf("Hola, %v. Welcome", nombre)
	return message
}
