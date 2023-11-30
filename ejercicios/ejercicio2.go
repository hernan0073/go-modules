package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ListadeNumeros() string {

	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {
		fmt.Println("Ingrese Numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, i*numero)
	}
}
