package main

import(
	"fmt"
	//"errors"
)

func main()  {
	var salary int = 16000
	if salary < 150000 {
        err := &myError{"el salario ingresado no alcanza el mínimo imponible"}
        println(err.Error())
		
    } else {
        fmt.Println("Debe pagar impuesto")
    }

}

type myError struct{
	msg string
}

func (e *myError) Error() string{
	return fmt.Sprintf("Error: %s", e.msg)
}