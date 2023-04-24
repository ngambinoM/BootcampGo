package main

import (
	"errors"
	"fmt"

)

func main()  {
	var salary int = 100
	e:= x()
	if salary <= 10000 {
		if errors.Is(e,err1) {
			fmt.Println(e.Error())
		}
		
    } else {
        fmt.Println("Cobra mas de 10.000")
    }

}

var err1 = errors.New("el salario es menor a 10.000")

func x() error{
	return fmt.Errorf("error: %w",err1)
}