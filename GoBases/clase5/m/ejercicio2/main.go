package main

import (
	"errors"
	"fmt"
	
)

func main()  {
	var salary int = 1600
	if salary <= 10000 {
        err := myError{"el salario es menor a 10.000"}
        if errors.Is(err,myError{"el salario es menor a 10.000"}){
			fmt.Println(err.Error())
		}
		
    } else {
        fmt.Println("Cobra mas de 10.000")
    }

}

type myError struct{
	msg string
}

func (e myError) Error() string{
	return e.msg
}

