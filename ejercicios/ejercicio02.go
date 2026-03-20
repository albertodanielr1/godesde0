package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() string {
	var numero int
	var err error
	var texto string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d X %d = %d \n", numero, i, i*numero)
	}

	return texto
}
