package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Apellido string
var Sueldo float32
var Fecha time.Time

func RestoVariable() {

	Nombre = "Hernan"
	Apellido = "Settino"
	Sueldo = 200000
	Fecha = time.Now()

	fmt.Println("Tu nombre es", Nombre)
	fmt.Println("Tu Apellido es", Apellido)
	fmt.Println("tu sueldo es", Sueldo)
	fmt.Println("La fecha es", Fecha)
}

func ConviertetoTex(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto

}
