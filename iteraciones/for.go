package iteraciones

import (
	"fmt"
)

func Iterar() {
	/*en el FOR , en i++ si cloco i+=5 me dara iteraciones de 5 en 5
	o depende del numero que coloque,
	igual con los opeadores si colo i-- seria cuenta regresiva*/
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
