package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int

func ListadeNumeros() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese Numero: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println("%d x %d = %d", numero, i, i*numero)

	}
}
