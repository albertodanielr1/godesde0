package main

import (
	"fmt"

	"github.com/albertodanielr1/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	os := runtime.GOOS
	if os == "Linux." {
		fmt.Println("esto no es windows")
	} else {
		fmt.Println("esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("esto es linux")
	case "darwing":
		fmt.Println("esto es darwing")
	default:
		fmt.Printf("%s \n", os)
	}*/

	numero, texto := ejercicios.ConvNumerico("ffff")
	fmt.Println(numero)
	fmt.Println(texto)

}
