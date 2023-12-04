package main

import (
	"fmt"
	"os"

	"github.com/hernan0073/go-modules/ejercicios"
)

func main() {
	var text string = ejercicios.ListadeNumeros()
	archivo, err := os.Create("./tabla.txt")
	if err != nil {
		fmt.Println("error al crear el archivo" + err.Error())
		return

	}
	fmt.Fprintln(archivo, text)
	archivo.Close()
}

// func SumaTabla() {
// 	var text string = ejercicio.ListadeNumeros()
// 	archivo, err := os.a("./files.txt/tabla.txt")
// 	if err != nil {
// 		fmt.Println("error al crear el archivo" + err.Error())
// 		return
// 	}
// 	fmt.Fprintln(archivo, text)
// 	archivo.Close()
// }
