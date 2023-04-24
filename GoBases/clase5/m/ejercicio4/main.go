package main

import (
	
	"fmt"
	
)

func main() {
	var salary int = 5000

	if salary < 150000 {
		err := fmt.Errorf("error: el minimo imponible es de 150.000",)
		panic(err)
	} else {
		fmt.Println("El salario es mayor a 10.000")
	}
}